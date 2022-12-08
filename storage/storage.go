package storage

import "UacademyGo/Blogpost/article_service/models"

type StorageInter interface {
	//* Article
	AddNewArticle(id string, box models.CreateModelArticle) error
	GetArticleById(id string) (models.GetByIDArticleModel, error)
	GetArticleList(offset, limit int, search string) (dataset []models.Article, err error)
	UpdateArticle(box models.UpdateArticleModel) error
	DeleteArticle(id string) error
	//* Author
	AddAuthor(id string, box models.CreateModelAuthor) error
	GetAuthorById(id string) (models.Author, error)
	GetAuthorList(limit, offset int, search string) (dataset []models.Author, err error)
	UpdateAuthor(box models.UpdateAuthorResponse) error
	DeleteAuthor(id string) error
}
