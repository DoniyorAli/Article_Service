package article

import (
	"UacademyGo/Blogpost/article_service/models"
	blogpost "UacademyGo/Blogpost/article_service/protogen/blogpost"
	"UacademyGo/Blogpost/article_service/storage"
	"context"
	"log"


	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type articleService struct {
	stg storage.StorageInter
	blogpost.UnimplementedArticleServiceServer
}

func NewArticleService(stg storage.StorageInter) *articleService {
	return &articleService{
		stg: stg,
	}
}

func (s *articleService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong {
		Message: "OK",
	}, nil
}

//?==============================================================================================================

func (s *articleService) CreateArticle(ctx context.Context, req *blogpost.CreateArticleRequest) (*blogpost.Article, error) {
	id := uuid.New()

	err := s.stg.AddNewArticle(id.String(), models.CreateModelArticle{
		Content: models.Content{
			Title: req.Content.Title,
			Body: req.Content.Body,
		},
		AuthorID: req.AuthorId,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.AddNewArticle: %s", err.Error())
	}

	article, err := s.stg.GetArticleById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleById: %s", err.Error())
	}

	var updatedAt string
	if article.UpdateAt != nil {
		updatedAt = article.UpdateAt.String()
	}

	return &blogpost.Article{
		Id: article.ID,
		Content: &blogpost.Content{
			Title: article.Title,
			Body: article.Body,
		},
		AuthorId: article.Author.ID,
		CreatedAt: article.CreateAt.String(),
		UpdatedAt: updatedAt,
	}, nil
}

//?==============================================================================================================

func (s *articleService) UpdateArticle(ctx context.Context, req *blogpost.UpdateArticleRequest) (*blogpost.Article, error) {
	err := s.stg.UpdateArticle(models.UpdateArticleModel{
		ID: req.Id,
		Content: models.Content{
			Title: req.Content.Title,
			Body: req.Content.Body,
		},
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateArticle: %s", err.Error())
	}

	article, err := s.stg.GetArticleById(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleById: %s", err.Error())
	}

	var updatedAt string
	if article.UpdateAt != nil {
		updatedAt = article.UpdateAt.String()
	}
	
	return &blogpost.Article{
		Id: article.ID,
		Content: &blogpost.Content{
			Title: article.Title,
			Body: article.Body,
		},
		AuthorId: article.Author.ID,
		CreatedAt: article.CreateAt.String(),
		UpdatedAt: updatedAt,
	}, nil 
}

//?==============================================================================================================

func (s *articleService) DeleteArticle(ctx context.Context, req *blogpost.DeleteArticleRequest) (*blogpost.Article, error) {

	article, err := s.stg.GetArticleById(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleById: %s", err.Error())
	}

	var updatedAt string
	if article.UpdateAt != nil {
		updatedAt = article.UpdateAt.String()
	}

	err = s.stg.DeleteArticle(article.ID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteArticle: %s", err.Error())
	}

	return &blogpost.Article{
		Id: article.ID,
		Content: &blogpost.Content{
			Title: article.Title,
			Body: article.Body,
		},
		AuthorId: article.Author.ID,
		CreatedAt: article.CreateAt.String(),
		UpdatedAt: updatedAt,
	}, nil 
}

//?==============================================================================================================

func (s *articleService) GetArticleList(ctx context.Context, req *blogpost.GetArticleListRequest) (*blogpost.GetArticleListResponse, error) {
	res := &blogpost.GetArticleListResponse{
		Articles: make([]*blogpost.Article, 0), //?????????????????????????
	}

	articleList, err := s.stg.GetArticleList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleList: %s", err.Error())
	}

	for _, v := range articleList {

		var updatedAt string
		if v.UpdateAt != nil {
			updatedAt = v.UpdateAt.String()
		}

		res.Articles = append(res.Articles, &blogpost.Article{
			Id: v.ID,
			Content: &blogpost.Content{
			Title: v.Title,
			Body: v.Body,
		},
		AuthorId: v.AuthorID,
		CreatedAt: v.CreateAt.String(),
		UpdatedAt: updatedAt,
		})
	}

	return res, nil
}

//?==============================================================================================================

func (s *articleService) GetArticleByID(ctx context.Context, req *blogpost.GetArticleByIDRequest) (*blogpost.GetArticleByIDResponse, error) {
	article, err := s.stg.GetArticleById(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleById: %s", err.Error())
	}

	if article.DeletedAt != nil {
		return nil, status.Errorf(codes.NotFound, "s.stg.GetArticleById: %s", err.Error())
	}

	var updatedAt string
	if article.UpdateAt != nil {
		updatedAt = article.UpdateAt.String()
	}

	var authorUpdatedAt string
	if article.Author.UpdateAt != nil {
		updatedAt = article.Author.UpdateAt.String()
	}
	
	return &blogpost.GetArticleByIDResponse{
		Id: article.ID,
		Content: &blogpost.Content{
			Title: article.Title,
			Body: article.Body,
		},
		Author: &blogpost.GetArticleByIDResponse_Author{
			Id: article.Author.ID,
			Fullname: article.Author.Fullname,
			Middlename: article.Author.Middlename,
			CreatedAt: article.CreateAt.String(),
			UpdatedAt: authorUpdatedAt,
		},
		CreatedAt: article.CreateAt.String(),
		UpdatedAt: updatedAt,
	}, nil
}