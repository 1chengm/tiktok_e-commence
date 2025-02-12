package rpc

import (
	productcatalogservice "gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"gomall/app/frontend/conf"
	frontendUtils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/user/userservice"
	cartservice "gomall/rpc_gen/kitex_gen/cart/cartservice"
	checkoutservice "gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	ProductClient productcatalogservice.Client
	CartClient cartservice.Client
	CheckoutClient checkoutservice.Client
	once sync.Once
)
func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
	})
}
func initUserClient() {
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
	
}
func initProductClient() {
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
	
}
func initCartClient() {
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
	
}
func initCheckoutClient(){
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}