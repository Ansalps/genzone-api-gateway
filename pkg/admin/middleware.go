package admin

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Ansalps/genzone-api-gateway/pkg/admin/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}
func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	fmt.Println("is it reaching in authrequired of admin")
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		fmt.Println("problem 1")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := strings.Split(authorization, "Bearer ")
	fmt.Println("is token splited")
	if len(token) < 2 {
		fmt.Println("problem 2")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	fmt.Println("before validate")
	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	fmt.Println("after validate")
	if err != nil {
		fmt.Println("print err!=nil",err)
	}
	// if res.Status != http.StatusOK {
	// 	fmt.Println("print res.Status != http.StatusOK")
	// }
	if err != nil || res.Status != http.StatusOK {
		fmt.Println("here apply")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("userId", res.UserId)
	ctx.Next()
}
