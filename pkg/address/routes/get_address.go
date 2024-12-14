package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ansalps/genzone-api-gateway/pkg/address/pb"
	"github.com/gin-gonic/gin"
)

func GetAddress(ctx *gin.Context, c pb.AddressServiceClient) {
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
	res, err := c.GetAddress(context.Background(), &pb.GetAddressRequest{
		Userid: userIdStr,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
