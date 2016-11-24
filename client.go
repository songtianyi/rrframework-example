package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/songtianyi/rrframework-example/proto/rrfp"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/rrframework/server"
	"github.com/songtianyi/rrframework/utils"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8003")
	if err != nil {
		logs.Error(err)
		return
	}
	c := rrserver.NewTCPConnection(conn)

	msg := new(rrfp.Message)
	msg.Hd = &rrfp.Head{
		rrutils.NewV4().String(),
		"rrfp.ExampleEchoRequest",
	}
	msg.By = &rrfp.Body{
		MsgType: &rrfp.Body_ExampleEchoRequest{
			ExampleEchoRequest: &rrfp.ExampleEchoRequest{Msg: "hello world!"},
		},
	}
	logs.Debug("before marshal:", msg)
	b, err := proto.Marshal(msg)
	if err != nil {
		logs.Error(err)
		return
	}

	if err := c.Write(b); err != nil {
		logs.Error(err)
		return
	}

	err, packet := c.Read()
	if err != nil {
		logs.Error(err)
		return
	}
	m := new(rrfp.Message)
	proto.Unmarshal(packet, m)
	logs.Info(m.String())
	logs.Debug("Response msg", m.GetBy().GetExampleEchoResponse().Msg)

}
