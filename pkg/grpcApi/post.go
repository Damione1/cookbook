package grpcApi

import (
	"context"
	"regexp"

	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/pkg/pbx"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (server *Server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateCreatePostRequest(req); err != nil {
		return nil, err
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

func validateUpdatePostRequest(req *pb.CreatePostRequest) error {
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

func validatePost(post *pb.Post) error {
	return validation.ValidateStruct(post,
		validation.Field(&post.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&post.Content, validation.Required),
		validation.Field(&post.Slug, validation.Length(1, 255), validation.Match(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`))),
		validation.Field(&post.Excerpt, validation.Length(1, 255)),
	)
}

func (server *Server) ListPosts(ctx context.Context, req *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	err := validateListPostsRequest(req)
	if err != nil {
		return nil, err
	}

	pageSize := int(req.GetPageSize())
	pageToken := int(req.GetPageToken())

	// Set default page size if not provided or if it's greater than the maximum allowed
	var maxPageSize int = 50
	if pageSize <= 0 || pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	offset := pageSize * pageToken

	dbPosts, err := models.Posts(
		qm.OrderBy("created_at desc"),
		qm.Limit(pageSize),
		qm.Offset(offset),
	).All(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	posts := make([]*pb.Post, 0, len(dbPosts))
	for _, dbPost := range dbPosts {
		posts = append(posts, pbx.DbPostToProto(dbPost))
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
		validation.Field(&req.PageSize, validation.Required, is.Int),
		validation.Field(&req.PageToken, validation.Required, is.Int),
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
		validation.Field(&req.Id, validation.Required, is.Int),
	)
}
