package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ansalps/genzone-api-gateway/pkg/auth/pb"
	"github.com/Ansalps/genzone-api-gateway/pkg/auth/utils"
	"github.com/Ansalps/genzone-api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context, c pb.AuthServiceClient) {
	body := models.SignUpRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the content of the JSON
	if err := utils.Validate(body); err != nil {
		fmt.Println("", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":     false,
			"message":    err.Error(),
			"error_code": http.StatusBadRequest,
		})
		return
	}
	if body.Password != body.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}
	res, err := c.SignUp(context.Background(), &pb.SignUpRequest{
		Firstname:       body.FirstName,
		Lastname:        body.LastName,
		Email:           body.Email,
		Password:        body.Password,
		Confirmpassword: body.ConfirmPassword,
		Phone:           body.Phone,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
