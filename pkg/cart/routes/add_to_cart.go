package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ansalps/genzone-api-gateway/pkg/cart/pb"
	"github.com/Ansalps/genzone-api-gateway/pkg/auth/utils"
	"github.com/Ansalps/genzone-api-gateway/pkg/models"
	"github.com/gin-gonic/gin"
)

func AddToCart(ctx *gin.Context, c pb.CartServiceClient) {
	b := models.CartRequestBody{}
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
	res, err := c.AddToCart(context.Background(), &pb.CreateCartRequest{
		Userid: userIdStr,
		Productid: b.ProductID,
		Quantity: int64(b.Quantiy),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
func GetCartItems(ctx *gin.Context, c pb.CartServiceClient) {
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
	fmt.Println("useridstr",userIdStr)
	res, err := c.GetCartItems(context.Background(), &pb.GetCartItemsRequest{
		Userid: userIdStr,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	// Check if cart items are empty and modify the response accordingly
    if len(res.Items) == 0 {
        ctx.JSON(http.StatusOK, gin.H{
            "status": res.Status,
            "message": "Cart is empty",
            "items":   []pb.CartItem{}, // Return an empty list for consistency
        })
        return
    }

    // If cart items exist, return the original response
    ctx.JSON(int(res.Status), gin.H{
        "status": res.Status,
        "message": "Cart items fetched successfully",
        "items":   res.Items,
    })
}