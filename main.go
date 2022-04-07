package main

import (
	com_anma "com.anma/com.anma"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net"
)

func main() {

	cm := com_anma.Comment{Body: "comment body"}
	msg := com_anma.Message{
		Body:    "some body",
		Comment: &cm,
	}

	fmt.Println(msg)
	WriteToFile("msg", &msg)
	var msg2 = com_anma.Message{}
	ReadFromFile("msg", &msg2)
	fmt.Println(msg2)

	//jsnMsg := ToJson()
	//fmt.Println(jsnMsg)
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return
	}
	serv := grpc.Server{}
	err = serv.Serve(listen)
	if err != nil {
		log.Println(err)
	}
}

func ReadFromFile(fn string, msg *com_anma.Message) error {
	bts, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	err2 := proto.Unmarshal(bts, msg)
	if err2 != nil {
		fmt.Println("Error")
	}
	return nil
}

func WriteToFile(file string, msg *com_anma.Message) error {
	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile(file, out, 0644)
	return nil
}

func ToJson(msg *proto.Message) string {
	marsh := jsonpb.Marshaler{}
	out, err := marsh.MarshalToString(&msg)
	if err != nil {
		return ""
	}
	return out
}

func fromJson(jsn string, pm proto.Message) {
	err := jsonpb.UnmarshalString(jsn, pm)
	if err != nil {
		log.Println(err)
	}
}
