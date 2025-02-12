package rpc

import (
	"gomall/app/checkout/conf"
	cartservice "gomall/rpc_gen/kitex_gen/cart/cartservice"
	paymentservice "gomall/rpc_gen/kitex_gen/payment/paymentservice"
	productcatalogservice "gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
	"github.com/cloudwego/kitex/client"
	checkoutUtils "gomall/app/checkout/utils"
	consul "github.com/kitex-contrib/registry-consul"
)
var(
	//cartservice
	//productservice
	//
	//paymentservice
	CartClient cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client

	once sync.Once
)
func InitClient(){
	once.Do(func(){
		initCartClient()
		initProductClient()
		initPaymentClient()
	})
}
func initCartClient(){
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	checkoutUtils.MustHandleError(err)
}
func initProductClient(){
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	checkoutUtils.MustHandleError(err)
}
func initPaymentClient(){
	// 服务发现
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	PaymentClient, err = paymentservice.NewClient("payment", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
	checkoutUtils.MustHandleError(err)
}