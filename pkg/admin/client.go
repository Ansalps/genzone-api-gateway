package admin

import (
	"fmt"

	"github.com/Ansalps/genzone-api-gateway/pkg/admin/pb"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	//usingWithInsecure()because of no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	//return the grpc client
	return pb.NewAuthServiceClient(cc)
}
