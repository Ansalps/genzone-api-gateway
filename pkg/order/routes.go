package order

import (
	"github.com/Ansalps/genzone-api-gateway/pkg/auth"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"github.com/Ansalps/genzone-api-gateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	//routes.Use(a.AuthRequired)
	r.POST("/order", a.AuthRequired, svc.CreateOrder)
	r.GET("/order",a.AuthRequired,svc.GetUserOrders)
	r.GET("/orderitems",a.AuthRequired,svc.GetOrderItems)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
func (svc *ServiceClient) GetUserOrders(ctx *gin.Context) {
	routes.GetUserOrders(ctx, svc.Client)
}
func (svc *ServiceClient) GetOrderItems(ctx *gin.Context) {
	routes.GetOrderItems(ctx, svc.Client)
}
