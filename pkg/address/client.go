package address

import (
	"fmt"

	"github.com/Ansalps/genzone-api-gateway/pkg/address/pb"
	"github.com/Ansalps/genzone-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AddressServiceClient
}

func InitServiceClient(c *config.Config) pb.AddressServiceClient {
	//using WithInsecure()because of no SSL running
	cc, err := grpc.Dial(c.AddressSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewAddressServiceClient(cc)
}
