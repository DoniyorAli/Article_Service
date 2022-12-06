package handlers

import (
	"net/http"
	"strconv"

	"UacademyGo/Blogpost/article_service/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// * ================== Create Author =========================
// CreateAuthor godoc
// @Summary     Create author
// @Description Create a new author
// @Tags        authors
// @Accept      json
// @Param       author body models.CreateModelAuthor true "author body"
// @Produce     json
// @Success     201 {object} models.JSONRespons{data=string}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/author [post]
func (h *handler) CreateAuthor(ctx *gin.Context) {
	var body models.CreateModelAuthor
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	id := uuid.New()

	err := h.Stg.AddAuthor(id.String(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	_, err = h.Stg.GetAuthorById(id.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONRespons{
		Message: "Author | GetList",
		Data:    id,
	})
}

// * ==================== Get Author By Id ====================
// GetAuthorById godoc
// @Summary     get author by id
// @Description get a new author
// @Tags        authors
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     404 {object} models.JSONErrorRespons
// @Router      /v1/author/{id} [get]
func (h *handler) GetAuthorById(ctx *gin.Context) {
	idStr := ctx.Param("id")

	//TODO UUID validation

	author, err := h.Stg.GetAuthorById(idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "passed successfully",
		Data:    author,
	})
}

// * ==================== Get Article List ====================
// GetArticleList godoc
// @Summary     List authors
// @Description get authors
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "10"
// @Param       search query    string false "smth"
// @Success     200 {object} models.JSONRespons{data=[]models.Author}
// @Router      /v1/author [get]
func (h *handler) GetAuthorList(ctx *gin.Context) {
	offsetStr := ctx.DefaultQuery("offset", "0")
	limitStr := ctx.DefaultQuery("limit", "10")
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

	authorList, err := h.Stg.GetAuthorList(offset, limit, searchStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "OK",
		Data:    authorList,
	})
}

// * ==================== Update Author =======================
// UpdateAuthor godoc
// @Summary     Update author
// @Description Update a new author
// @Tags        authors
// @Accept      json
// @Param       author body models.UpdateAuthorResponse true "updating author"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/author [put]
func (h *handler) UpdateAuthor(ctx *gin.Context) {
	var body models.UpdateAuthorResponse
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	err := h.Stg.UpdateAuthor(body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Message: "Storage error",
			Error:   err.Error(),
		})
		return
	}

	author, err := h.Stg.GetAuthorById(body.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "Author successfully updated",
		Data:    author,
	})
}

// * ==================== Delete Author =======================
// DeleteAuthor godoc
// @Summary     Delete author
// @Description delete author
// @Tags        authors
// @Accept      json
// @Param       id path string true "Author ID"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v1/author/{id} [delete]
func (h *handler) DeleteAuthor(ctx *gin.Context) {
	idStr := ctx.Param("id")

	author, err := h.Stg.GetAuthorById(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorRespons{
			Message: "author already deleted or not found or you entered wrong ID",
			Error:   err.Error(),
		})
		return
	}

	err = h.Stg.DeleteAuthor(author.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNotFound, models.JSONRespons{
		Message: "Author suucessfully deleted",
		Data:    author,
	})
}
