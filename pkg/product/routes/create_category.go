package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ansalps/genzone-api-gateway/pkg/auth/utils"
	"github.com/Ansalps/genzone-api-gateway/pkg/models"
	"github.com/Ansalps/genzone-api-gateway/pkg/product/pb"
	"github.com/gin-gonic/gin"
)

func CreateCategory(ctx *gin.Context, c pb.ProductServiceClient) {
	fmt.Println("hi on category")
	body := models.CategoryRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	//validate the content of the json
	err := utils.Validate(body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    gin.H{},
		})
		return
	}
	res, err := c.CreateCategory(context.Background(), &pb.CreateCategoryRequest{
		Categoryname: body.CategoryName,
		Description:  body.Description,
		Imageurl:     body.ImageUrl,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}

func ListCategories(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.ListCategories(context.Background(), &pb.EmptyRequest{
		
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}