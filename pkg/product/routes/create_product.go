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

type product struct {
	Id           int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Categoryname string  `protobuf:"bytes,2,opt,name=categoryname,proto3" json:"categoryname,omitempty"`
	Productname  string  `protobuf:"bytes,3,opt,name=productname,proto3" json:"productname,omitempty"`
	Description  string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Imageurl     string  `protobuf:"bytes,5,opt,name=imageurl,proto3" json:"imageurl,omitempty"`
	Price        float64 `protobuf:"fixed64,6,opt,name=price,proto3" json:"price"`
	Stock        int64   `protobuf:"varint,7,opt,name=stock,proto3" json:"stock"`
	Popular      bool    `protobuf:"varint,8,opt,name=popular,proto3" json:"popular"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	fmt.Println("hi on product")
	body := models.ProductRequestBody{}
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
	fmt.Println("body.Price", body.Price)
	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Categoryname: body.CategoryName,
		Productname:  body.ProductName,
		Description:  body.Description,
		Imageurl:     body.ImageUrl,
		Price:        body.Price,
		Stock:        int64(body.Stock),
		Popular:      body.Popular,
		Size:         body.Size,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
func ListProducts(ctx *gin.Context, c pb.ProductServiceClient) {
	res, err := c.ListProducts(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	// Initialize products slice
	var products []product

	// Map Protobuf response to new product struct
	for _, item := range res.Products {
		// Create a new product for each iteration
		newProduct := product{
			Id:           item.Id,
			Categoryname: item.Categoryname,
			Productname:  item.Productname,
			Description:  item.Description,
			Imageurl:     item.Imageurl,
			Price:        item.Price,
			Stock:        item.Stock,
			Popular:      item.Popular,
		}
		products = append(products, newProduct)
	}
	response := gin.H{
		"products": products,
	}

	if res.Error != "" { // Include the "error" tag only if res.Error is not empty
		response["error"] = res.Error
	}

	ctx.JSON(int(res.Status), response)
}
