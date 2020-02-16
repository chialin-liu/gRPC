package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	simplepb "github.com/simple"
)

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cant convert to JSON", err)
		return ""
	}
	return out
}
func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Fail convert to pb", err)
	}
}
func main() {
	sm := doSimple()
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
	smAsString := toJSON(sm)
	fmt.Println(smAsString)
	sm3 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm3)
	fmt.Println(sm3)
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Read Fail", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Could not read data to structs", err2)
		return err2
	}
	return nil
}
func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cant serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Cant write to file", err)
		return err
	}
	fmt.Println("Date write to file successful")
	return nil
}
func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple",
		SampleList: []int32{1, 4, 5},
	}
	sm.Name = "Yorick"
	fmt.Println(sm.GetId())
	fmt.Println(sm)
	return &sm
}
