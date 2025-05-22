package controllers

import (
	"goTask/internal/DTO"
	"goTask/internal/generationService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PasswordGenerationController struct {
	service generationService.PasswordGenerationService
}

func NewPasswordGenerationController(service generationService.PasswordGenerationService) *PasswordGenerationController {
	return &PasswordGenerationController{
		service: service,
	}
}

func (c *PasswordGenerationController) Generate(ctx *gin.Context) {
	var data DTO.GenerationOptions

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusOK, getResponse(
			map[string]any{"errors": []string{"Invalid input data (length must be a number, and at least one option must be selected)"}},
			false,
			"Something went wrong.",
		))

		return
	}

	password, errs := c.service.Generate(data)
	if len(errs) > 0 {
		ctx.JSON(http.StatusOK, getResponse(
			map[string]any{"errors": errs},
			false,
			"Something went wrong.",
		))

		return
	}

	ctx.JSON(http.StatusOK, getResponse(
		map[string]any{"password": password},
		true,
		"",
	))
}

func (c *PasswordGenerationController) GetGenerationOptions(ctx *gin.Context) {
	options := c.service.GetGenerationOptions()

	if options == nil {
		ctx.JSON(
			http.StatusOK,
			getResponse(
				nil,
				false,
				"Something went wrong during the generation options request. Please try again later.",
			),
		)

		return
	}

	ctx.JSON(
		http.StatusOK,
		getResponse(
			map[string]any{
				"generationOptions": options,
			},
			true,
			"",
		),
	)
}

func getResponse(data map[string]any, isSuccess bool, message string) gin.H {
	return gin.H{
		"data":      data,
		"isSuccess": isSuccess,
		"message":   message,
	}
}
