package rpc

import (
	"sync"
	
	productcatalogservice "gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"gomall/app/cart/conf"
	cartUtils "gomall/app/cart/utils"
	
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once sync.Once
)
func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	cartUtils.MustHandleError(err)
	
}