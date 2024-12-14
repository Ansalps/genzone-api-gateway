package auth

import (
	"fmt"

	"github.com/Ansalps/genzone-api-gateway/pkg/auth/routes"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	r.POST("/signup", svc.SignUp)
	r.POST("/login", svc.Login)
	return svc
}
func (svc *ServiceClient) SignUp(ctx *gin.Context) {
	fmt.Println("second hi")
	routes.SignUp(ctx, svc.Client)
}
func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
