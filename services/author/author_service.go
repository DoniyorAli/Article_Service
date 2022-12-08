package author

import (
	blogpost "UacademyGo/Blogpost/article_service/protogen/blogpost"
	"context"
	"log"
)

type AuthorService struct {
	blogpost.UnimplementedAuthorServiceServer
}

func (s *AuthorService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong {
		Message: "OK",
	}, nil
}