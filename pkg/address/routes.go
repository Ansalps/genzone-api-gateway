package address

import (
	"github.com/Ansalps/genzone-api-gateway/pkg/address/routes"
	"github.com/Ansalps/genzone-api-gateway/pkg/auth"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routes := r.Group("/address")
	//routes.Use(a.AuthRequired)
	routes.POST("/", a.AuthRequired, svc.CreateAddress)
	routes.POST("/getaddress", a.AuthRequired, svc.GetAddress)
}
func (svc *ServiceClient) CreateAddress(ctx *gin.Context) {
	routes.CreateAddress(ctx, svc.Client)
}
func (svc *ServiceClient) GetAddress(ctx *gin.Context) {
	routes.GetAddress(ctx, svc.Client)
}
