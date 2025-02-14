package email

import (
	"gomall/app/email/infra/mq"
	"gomall/app/email/infra/notify"
	"gomall/rpc_gen/kitex_gen/email"

	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ConsumerInit() {

	sub, err := mq.Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq //定义消息结构体
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Errorf("unmarshal email req failed: %v", err)
			return 

		}
		noopEmail := notify.NewNoopEmail()
		_=noopEmail.Send(&req) //调用发送邮件的方法
	})
	if err != nil {
		klog.Errorf("subscribe email failed: %v", err)
		return 
	}//订阅消息失败，打印错误信息

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		mq.Nc.Close()
	}) //关闭订阅
}