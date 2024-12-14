package cart

import (
	"github.com/Ansalps/genzone-api-gateway/pkg/auth"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"github.com/Ansalps/genzone-api-gateway/pkg/cart/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	//routes.Use(a.AuthRequired)
	r.POST("/cart", a.AuthRequired, svc.AddToCart)
	r.GET("/cart",a.AuthRequired,svc.GetCartItems)
}

func (svc *ServiceClient) AddToCart(ctx *gin.Context) {
	routes.AddToCart(ctx, svc.Client)
}
func (svc *ServiceClient) GetCartItems(ctx *gin.Context) {
	routes.GetCartItems(ctx, svc.Client)
}
