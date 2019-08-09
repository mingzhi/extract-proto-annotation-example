package main

import (
	"fmt"
	"io/ioutil"
	"log"

	apb "github.com/mingzhi/extract-proto-annotation-example/google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	m := getFileDescriptor().Services().Get(0).Methods().Get(0)
	opts := m.Options().(*descriptorpb.MethodOptions)
	if proto.HasExtension(opts, apb.E_Http) {
		v := proto.GetExtension(opts, apb.E_Http)
		r := v.(proto.Message).(*apb.HttpRule)
		fmt.Println(r)
	}
}

func getFileDescriptor() protoreflect.FileDescriptor {
	filename := "http_bookstore.protoset"
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error when reading %s: %v", filename, err)
	}
	protoset := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(bs, protoset); err != nil {
		log.Fatalf("Unable to parse %T from %s: %v", protoset, filename, err)
	}
	f, err := protodesc.NewFile(protoset.GetFile()[0], nil)
	if err != nil {
		log.Fatalf("protodesc.NewFile() error: %v", err)
	}
	return f
}
