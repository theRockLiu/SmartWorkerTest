// test_pb project main.go
package main

import (
	proto "code.google.com/p/goprotobuf/proto"
	"fmt"
	"log"
	"net"
	"SmartWorkerTest/myproto"
	"unsafe"
)

func main() {
	fmt.Println("Hello World!")

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	req := &myproto.RegReq{
		Label: proto.String("hello"),
		Type:  proto.Int32(17)}

	var buf [4]byte
	//*((*int)(unsafe.Pointer(&buf))) = proto.Size(req)

	type SMsgHeader struct {
		len uint32
	}

	//var smh SMsgHeader

	//const CONST_MSG_HEADER_LEN = unsafe.Sizeof(smh)
	msgHdr := (*SMsgHeader)(unsafe.Pointer(&buf))

	bytes, err := proto.Marshal(req)
	fmt.Println(proto.Size(req))
	fmt.Println(len(bytes))
	msgHdr.len = (uint32)(proto.Size(req))

	conn.Write(buf[:])
	conn.Write(bytes)

	ch := make(chan int)
	<-ch

}
