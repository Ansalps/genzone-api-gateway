package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ansalps/genzone-api-gateway/pkg/address/pb"
	"github.com/Ansalps/genzone-api-gateway/pkg/auth/utils"
	"github.com/Ansalps/genzone-api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreateAddress(ctx *gin.Context, c pb.AddressServiceClient) {
	b := models.AddressRequestBody{}
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
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId not found in context"})
		return
	}
	var userIdStr string
	switch v := userId.(type) {
	case string:
		userIdStr = v
	case int64:
		userIdStr = fmt.Sprintf("%d", v) // Convert int64 to string
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userId type"})
		return
	}
	fmt.Printf("userId type: %T, value: %v\n", userId, userId)
	res, err := c.CreateAddress(context.Background(), &pb.CreateAddressRequest{
		Userid:     userIdStr,
		Country:    b.Country,
		State:      b.State,
		District:   b.District,
		Streetname: b.StreetName,
		Pincode:    b.PinCode,
		Phone:      b.Phone,
		Default:    b.Default,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
