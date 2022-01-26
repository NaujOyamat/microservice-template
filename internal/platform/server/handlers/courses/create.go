package courses

import (
	"net/http"

	"Github.com/NaujOyamat/microservice-template/internal/domain/courses"
	"Github.com/NaujOyamat/microservice-template/internal/domain/courses/repository"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(repo repository.ICourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := courses.NewCourse(req.ID, req.Name, req.Duration)
		if err := repo.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
