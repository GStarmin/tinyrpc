// Code generated by protoc-gen-tinyrpc.

package message

import "errors"

// svr 文件：定义服务接口和方法，生成客户端存根代码和服务器框架代码，使得开发基于 Protobuf 和 RPC 的服务变得更加容易。

// ArithService Defining Computational Digital Services
type ArithService struct{}

// Add addition
func (*ArithService) Add(args *ArithRequest, reply *ArithResponse) error {
	reply.C = args.A + args.B
	return nil
}

// Sub subtraction
func (*ArithService) Sub(args *ArithRequest, reply *ArithResponse) error {
	reply.C = args.A - args.B
	return nil
}

// Mul multiplication
func (*ArithService) Mul(args *ArithRequest, reply *ArithResponse) error {
	reply.C = args.A * args.B
	return nil
}

// Div division
func (*ArithService) Div(args *ArithRequest, reply *ArithResponse) error {
	if args.B == 0 {
		return errors.New("divided is zero")
	}
	reply.C = args.A / args.B
	return nil
}
