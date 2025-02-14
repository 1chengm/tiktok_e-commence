package notify

import (
	"gomall/rpc_gen/kitex_gen/email"

	"github.com/kr/pretty"
)
type NoopEmail struct{}


func (e *NoopEmail) Send(req *email.EmailReq) error {
	pretty.Printf("send email: %+v\n", req)
	return nil
}
func NewNoopEmail() NoopEmail{

	return NoopEmail{}
}