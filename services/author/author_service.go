package author

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

type authorService struct {
	stg storage.StorageInter
	blogpost.UnimplementedAuthorServiceServer
}

func NewAuthorService(stg storage.StorageInter) *authorService {
	return &authorService{
		stg: stg,
	}
}

func (s *authorService) Ping(ctx context.Context, req *blogpost.Empty) (*blogpost.Pong, error) {
	log.Println("Ping")

	return &blogpost.Pong {
		Message: "OK",
	}, nil
}

//?==============================================================================================================

func (s *authorService) CreateAuthor(ctx context.Context, req *blogpost.CreateAuthorRequest) (*blogpost.CreateAuthorResponse, error) {
	id := uuid.New()

	err := s.stg.AddAuthor(id.String(), models.CreateModelAuthor{
		Fullname: req.Fullname,
		Middlename: req.Middlename,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.CreateAuthor: %s", err.Error())
	}

	res, err := s.stg.GetAuthorById(id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleById: %s", err.Error())
	}

	return &blogpost.CreateAuthorResponse{
		Id:	res.ID,
	}, nil
}

func (s *authorService) GetAuthorByID(ctx context.Context, req *blogpost.GetAuthorByIDRequest) (*blogpost.Author, error) {
	author, err := s.stg.GetAuthorById(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetAuthorById: %s", err.Error())
	}
	var updated_at string
	if author.UpdateAt !=nil{
		updated_at=author.UpdateAt.String()
	}
	return &blogpost.Author{
		Id:	author.ID,
		Fullname:author.Fullname,
		Middlename: author.Middlename,
		CreatedAt: author.CreateAt.String(),
		UpdatedAt: updated_at,
	}, nil
}


func (s *authorService) GetAuthorList(ctx context.Context, req *blogpost.GetAuthorListRequest) (*blogpost.GetAuthorListResponse, error) {
	res := &blogpost.GetAuthorListResponse{
		Authors: make([]*blogpost.Author, 0),
	}

	articleList, err := s.stg.GetAuthorList(int(req.Offset), int(req.Limit), req.Search)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.GetArticleList: %s", err.Error())
	}
	
	for _, v := range articleList {

		var updatedAt string 
		if v.UpdateAt != nil {
			updatedAt = v.UpdateAt.String()
		}

		res.Authors = append(res.Authors, &blogpost.Author{
			Id: v.ID,
			Fullname: v.Fullname,
			Middlename: v.Middlename,
			CreatedAt: v.CreateAt.String(),
			UpdatedAt: updatedAt,
		})
	}

	return res, nil
}




func (s *authorService) UpdateAuthor(ctx context.Context, req *blogpost.UpdateAuthorRequest) (*blogpost.UpdateAuthorResponse, error) {

	err := s.stg.UpdateAuthor(models.UpdateAuthorResponse{
		ID: req.Id,
		Fullname: req.Fullname,
		Middlename: req.Middlename,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateAuthor: %s", err.Error())

	}

	_, err = s.stg.GetAuthorById(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.UpdateAuthor: %s", err.Error())
	}

	return &blogpost.UpdateAuthorResponse{
		Status: "Updated",
	}, nil
}



func (s *authorService) DeleteAuthor(ctx context.Context, req *blogpost.DeleteAuthorRequest) (*blogpost.DeleteAuthorResponse, error) {

	err := s.stg.DeleteAuthor(req.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "s.stg.DeleteAuthor: %s", err.Error())

	}

	return &blogpost.DeleteAuthorResponse{
		Status: "Deleted",
	}, nil
}