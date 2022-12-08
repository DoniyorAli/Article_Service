package article

import (
	blogpost "UacademyGo/Blogpost/article_service/protogen/blogpost"
	"context"
	"log"
)

type ArticleService struct {
	blogpost.UnimplementedArticleServiceServer
}

func (s *ArticleService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong {
		Message: "OK",
	}, nil
}