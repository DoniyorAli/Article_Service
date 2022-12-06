package handlers

import (
	"net/http"
	"strconv"

	"UacademyGo/Article/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// * ================== Create Article ======================
// CreateArticle godoc
// @Summary     Create article
// @Description Create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body models.CreateModelArticle true "article body"
// @Success     201 {object} models.JSONRespons{data=models.Article} 
// @Failure     400 {object} models.JSONErrorRespons                 
// @Router      /v1/article [post]
func (h *handler) CreateArticle(ctx *gin.Context) {
	var body models.CreateModelArticle
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	id := uuid.New()

	err := h.Stg.AddNewArticle(id.String(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleById(id.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONRespons{
		Message: "Article successfully created",
		Data:    article,
	})
}

// * ==================== GetArticleById ====================
// GetArticleById godoc
// @Summary     get article by id
// @Description get a new article
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.GetByIDArticleModel}
// @Failure     404 {object} models.JSONErrorRespons
// @Router      /v1/article/{id} [get]
func (h *handler) GetArticleById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	//TODO UUID validation

	article, err := h.Stg.GetArticleById(idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "passed successfully",
		Data:    article,
	})
}

// * ==================== GetArticleList ====================
// GetArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} models.JSONRespons{data=[]models.Article}
// @Router      /v1/article [get]
func (h *handler) GetArticleList(ctx *gin.Context) {
	
	offsetStr := ctx.DefaultQuery("offset", h.Cfg.DefaultOffset)
	limitStr := ctx.DefaultQuery("limit", h.Cfg.DefaultLimit)
	searchStr := ctx.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	articleList, err := h.Stg.GetArticleList(offset, limit, searchStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "OK",
		Data:    articleList,
	})
}

// * ==================== UpdateArticle ====================
// UpdateArticle godoc
// @Summary     Update article
// @Description Update a new article
// @Tags        articles
// @Accept      json
// @Param       article body models.UpdateArticleResponse true "updating article"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Article}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/article [put]
func (h *handler) UpdateArticle(ctx *gin.Context) {
	var body models.UpdateArticleResponse
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	err := h.Stg.UpdateArticle(body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleById(body.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "Article successfully updated",
		Data:    article,
	})
}

// * ==================== DeleteArticle ====================
// DeleteArticle godoc
// @Summary     Delete article
// @Description delete article
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Article}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/article/{id} [delete]
func (h *handler) DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")

	article, err := h.Stg.GetArticleById(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Message: "article already deleted or not found or you entered wrong ID",
			Error:   err.Error(),
		})
		return
	}

	err = h.Stg.DeleteArticle(article.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNotFound, models.JSONRespons{
		Message: "Article suucessfully deleted",
		Data:    article,
	})
}

// * ==================== PingPong ====================
func (h *handler) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
