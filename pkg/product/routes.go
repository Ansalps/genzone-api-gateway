package product

import (
	"github.com/Ansalps/genzone-api-gateway/pkg/admin"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"github.com/Ansalps/genzone-api-gateway/pkg/product/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *admin.ServiceClient) {
	a := admin.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	//routes.Use(a.AuthRequired)
	r.POST("/admin/category", a.AuthRequired, svc.CreateCategory)
	r.GET("/admin/category", a.AuthRequired, svc.ListCategories)
	r.POST("/admin/product", a.AuthRequired, svc.CreateProduct)
	r.GET("/admin/product", a.AuthRequired, svc.ListProducts)
}

func (svc *ServiceClient) CreateCategory(ctx *gin.Context) {
	routes.CreateCategory(ctx, svc.Client)
}
func (svc *ServiceClient) ListCategories(ctx *gin.Context) {
	routes.ListCategories(ctx, svc.Client)
}
func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
func (svc *ServiceClient) ListProducts(ctx *gin.Context) {
	routes.ListProducts(ctx, svc.Client)
}
