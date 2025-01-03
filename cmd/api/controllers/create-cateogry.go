package controllers

import (
	"github.com/gin-gonic/gin"
	use_cases "github.com/jose.dev2301/go-categories-microsservice/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding: "required"`
}

func CreateCategory(ctx *gin.Engine) {
	useCase := use_cases.NewCreateCategoryUseCase()

	useCase.Execute()
}
