package pbx

import (
	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/util"
	"github.com/volatiletech/null/v8"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DbPostToProto(post *models.Post) *pb.Post {
	postPb := &pb.Post{
		Id:         int32(post.ID),
		Title:      post.Title,
		Slug:       post.Slug,
		Content:    post.Content.String,
		Excerpt:    post.Excerpt.String,
		CreateTime: timestamppb.New(post.CreatedAt.Time),
		UpdateTime: timestamppb.New(post.UpdatedAt.Time),
	}
	return postPb
}

func ProtoPostToDb(post *pb.Post) *models.Post {
	postDb := &models.Post{
		Title:   post.GetTitle(),
		Content: null.NewString(post.GetContent(), post.GetContent() != ""),
	}
	if post.GetId() != 0 {
		postDb.ID = int(post.GetId())
	}
	if post.GetExcerpt() != "" {
		postDb.Excerpt = null.StringFrom(post.GetExcerpt())
	}
	if post.GetCreateTime() != nil {
		postDb.CreatedAt = null.TimeFrom(post.GetCreateTime().AsTime())
	}
	if post.GetUpdateTime() != nil {
		postDb.UpdatedAt = null.TimeFrom(post.GetUpdateTime().AsTime())
	}
	postDb.Slug = util.Slugify(post.GetTitle())
	if post.GetSlug() != "" {
		postDb.Slug = post.GetSlug()
	}

	return postDb
}
