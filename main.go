package main

import (
	"fmt"
	v1 "github.com/clutonik/elastic-cloud-tools/pkg/api/v1"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"log"
)

func main(){
	cluster := v1.Cluster{
		ClusterAddress:     "",
		UserName:           "admin",
		Password:           "admin",
		Exists:             false,
		DeploymentTemplate: v1.Cluster_HOT_WARM,
		CreationDate:       nil,
		SupportEmails:      []string{"email1", "email2"},
	}

	jsString, err := toJSON(&cluster)
	if err != nil{
		log.Fatalln("could not convert to JSON", err)
	}
	fmt.Println(jsString)

}

func toJSON(m proto.Message) (string, error){
	marshaler := jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       " ",
		AnyResolver:  nil,
	}
	out, err := marshaler.MarshalToString(m)
	if err != nil {
		return "", err
	}
	return out, nil
}