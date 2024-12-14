package main

import (
	"log"

	"github.com/Ansalps/genzone-api-gateway/pkg/address"
	"github.com/Ansalps/genzone-api-gateway/pkg/admin"
	"github.com/Ansalps/genzone-api-gateway/pkg/auth"
	"github.com/Ansalps/genzone-api-gateway/pkg/cart"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"github.com/Ansalps/genzone-api-gateway/pkg/order"
	"github.com/Ansalps/genzone-api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()
	authSvc := *auth.RegisterRoutes(r, &c)
	adminSvc := *admin.RegisterRoutes(r, &c)
	//admin.RegisterRoutes(r,&c)
	address.RegisterRoutes(r, &c, &authSvc)
	product.RegisterRoutes(r, &c, &adminSvc)
	cart.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	r.Run(c.Port)
}
