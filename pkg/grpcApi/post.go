package grpcApi

import (
	"context"
	"regexp"

	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/pkg/pbx"
	"github.com/friendsofgo/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	authorizeUserPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateCreatePostRequest(req); err != nil {
		return nil, err
	}

	user, err := models.FindUser(ctx, server.config.DB, authorizeUserPayload.ID.String())
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, errors.Wrap(err, "Unsufficient permissions to create a post")
	}

	post := pbx.ProtoPostToDb(req.GetPost())

	err = post.Insert(ctx, server.config.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &pb.CreatePostResponse{
		Post: pbx.DbPostToProto(post),
	}, nil
}

func validateCreatePostRequest(req *pb.CreatePostRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Post,
			validation.Required,
			validation.By(
				func(value interface{}) error {
					return validatePost(value.(*pb.Post))
				},
			),
		),
	)
}

func (server *Server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	if _, err := server.authorizeUser(ctx); err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateUpdatePostRequest(req); err != nil {
		return nil, err
	}

	post := pbx.ProtoPostToDb(req.GetPost())

	_, err := post.Update(ctx, server.config.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePostResponse{
		Post: pbx.DbPostToProto(post),
	}, nil
}

func validateUpdatePostRequest(req *pb.UpdatePostRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Post,
			validation.Required,
			validation.By(
				func(value interface{}) error {
					post := value.(*pb.Post)
					return validation.ValidateStruct(post, validation.Field(&post.Id, validation.Required))
				},
			),
			validation.By(
				func(value interface{}) error {
					return validatePost(value.(*pb.Post))
				},
			),
		),
	)
}

func (server *Server) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	err := validateListPostsRequest(req)
	if err != nil {
		return nil, err
	}

	pageSize := int(req.GetPageSize())
	pageToken := 1
	if req.GetPageToken() != 0 {
		pageToken = int(req.GetPageToken())
	}
	offset := (pageToken - 1) * pageSize

	dbPosts, err := models.Posts(
		qm.OrderBy("created_at desc"),
		qm.Limit(pageSize),
		qm.Offset(offset),
	).All(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	posts := make([]*pb.Post, len(dbPosts))
	for i, dbPost := range dbPosts {
		posts[i] = &pb.Post{
			Id:         int32(dbPost.ID),
			Title:      dbPost.Title,
			Slug:       dbPost.Slug,
			Excerpt:    dbPost.Excerpt.String,
			CreateTime: timestamppb.New(dbPost.CreatedAt.Time),
			UpdateTime: timestamppb.New(dbPost.UpdatedAt.Time),
		}
	}

	nextPageToken := 0
	if len(dbPosts) == pageSize {
		nextPageToken = pageToken + 1
	}

	return &pb.ListPostsResponse{
		Posts:         posts,
		NextPageToken: int32(nextPageToken),
	}, nil
}

func validateListPostsRequest(req *pb.ListPostsRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.PageSize, validation.Required, validation.Min(1), validation.Max(100)),
		validation.Field(&req.PageToken, validation.Min(1)),
	)
}

func (server *Server) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	err := validateGetPostRequest(req)
	if err != nil {
		return nil, err
	}

	dbPost, err := models.Posts(
		models.PostWhere.ID.EQ(int(req.GetId())),
	).One(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	return &pb.GetPostResponse{
		Post: pbx.DbPostToProto(dbPost),
	}, nil
}

func validateGetPostRequest(req *pb.GetPostRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required),
	)
}

func (server *Server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	if _, err := server.authorizeUser(ctx); err != nil {
		return nil, unauthenticatedError(err)
	}

	err := validateDeletePostRequest(req)
	if err != nil {
		return nil, err
	}

	_, err = models.Posts(
		models.PostWhere.ID.EQ(int(req.GetId())),
	).DeleteAll(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePostResponse{}, nil
}

func validateDeletePostRequest(req *pb.DeletePostRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required, is.Int),
	)
}

func validatePost(post *pb.Post) error {
	return validation.ValidateStruct(post,
		validation.Field(&post.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&post.Content, validation.Required),
		validation.Field(&post.Slug, validation.Length(1, 255), validation.Match(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`))),
		validation.Field(&post.Excerpt, validation.Length(1, 255)),
	)
}
