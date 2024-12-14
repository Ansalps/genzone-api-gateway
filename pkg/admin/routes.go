package admin

import (
	"github.com/Ansalps/genzone-api-gateway/pkg/admin/routes"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
	
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	r.POST("/admin/login", svc.Login)
	return svc
}
func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
