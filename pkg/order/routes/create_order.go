package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ansalps/genzone-api-gateway/pkg/auth/utils"
	"github.com/Ansalps/genzone-api-gateway/pkg/models"
	"github.com/Ansalps/genzone-api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	b := models.OrderRequestBody{}
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
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		Userid:    userIdStr,
		Addressid: b.AddressID,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}

func GetUserOrders(ctx *gin.Context, c pb.OrderServiceClient) {
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
        userIdStr = fmt.Sprintf("%d", v)
    default:
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid userId type"})
        return
    }

    fmt.Printf("userId type: %T, value: %v\n", userId, userId)
    fmt.Println("userIdStr", userIdStr)

    res, err := c.GetUserOrders(context.Background(), &pb.GetUserOrdersRequest{
        Userid: userIdStr,
    })
    if err != nil {
        ctx.AbortWithError(http.StatusBadGateway, err)
        return
    }

    // Log the full response for debugging
    fmt.Printf("GetUserOrders Response: %+v\n", res)

    if len(res.Orders) == 0 {
        ctx.JSON(http.StatusOK, gin.H{"message": "No orders found for this user", "status": res.Status})
        return
    }

    ctx.JSON(int(res.Status), gin.H{
        "status": res.Status,
        "orders": res.Orders,
    })
}


func GetOrderItems(ctx *gin.Context, c pb.OrderServiceClient) {
	b := models.OrderItemRequestBody{}
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
	fmt.Println("useridstr",userIdStr)
	res, err := c.GetOrderItems(context.Background(), &pb.GetOrderItemsRequest{
		Userid: userIdStr,
		Orderid: b.OrderID,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}

