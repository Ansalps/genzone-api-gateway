package routes

import (
	"context"
	"net/http"

	"github.com/Ansalps/genzone-api-gateway/pkg/auth/pb"
	"github.com/Ansalps/genzone-api-gateway/pkg/auth/utils"
	"github.com/Ansalps/genzone-api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := models.LoginRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//validate the content of the json
	err := utils.Validate(b)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    gin.H{},
		})
		return
	}
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
