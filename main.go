package main

import (
	com_anma "com.anma/src/com.anma"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {

	cm := com_anma.Comment{Body: "comment body"}
	msg := com_anma.Message{
		Body:    "some body",
		Comment: &cm,
	}

	fmt.Println(msg)
}

func ReadFromFile(fn string, msg *com_anma.Message) error {
	file, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	err2 := proto.Unmarshal(file, msg)
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
