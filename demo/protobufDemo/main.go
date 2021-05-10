package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"trueabc.top/zinx/demo/protobufDemo/pb"
)

func main() {
	p := &pb.Person{
		Name:   "TrueAbc",
		Age:    16,
		Emails: []string{"212@qq.com", "13ease"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "122222",
				Type:   pb.PhoneType_MOBILE,
			},
		},
	}

	a, _ := p.Marshal()
	fmt.Println(a)

	newP := &pb.Person{}
	proto.Unmarshal(a, newP)
	fmt.Println(newP, p)

}
