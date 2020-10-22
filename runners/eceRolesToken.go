package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	eceCoordinatorHost string
	eceRunnerRole string
	ecePassword string
	eceUser string
)

const (
	TOKEN_ENDPOINT = "/api/v1/platform/configuration/security/enrollment-tokens"
	ECE_PORT = "12400"
	HTTP_PROTOCOL = "http://"
)

type RoleTokenResponse struct {
	Token string `json:"token"`
}

func main(){
	// Command line flags
	flag.StringVar(&eceCoordinatorHost, "coordinator-host", os.Getenv("ECE_COORDINATOR_HOST"),"Specify ECE Coordinator hostname")
	flag.StringVar(&eceRunnerRole, "runner-role", "allocator", "Specify role type to generate Role Token for")
	flag.StringVar(&eceUser, "ece-user",os.Getenv("ECE_USER") , "Specify ECE root user")
	flag.StringVar(&ecePassword, "ece-password",os.Getenv("ECE_PASSWORD") , "Specify ECE Password")
	flag.Parse()
	if eceCoordinatorHost == "" || ecePassword == "" || eceUser == "" {
		fmt.Println("Command Usage: ")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var roles []string
	roles = append(roles, eceRunnerRole)
	roleToken, err := GetRolesToken(roles)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// Print roleToken
	fmt.Println(roleToken.Token)
}

// getRolesToken accepts slice of roles and returns a Role token for that role.
func GetRolesToken(roles []string) (*RoleTokenResponse, error){
	// Prepare payload for POST Request
	postBody := map[string]interface{}{
		"persistent": false,
		"roles": roles,
	}

	// Convert Go DataStructure (Map) to JSON
	jsonPayload, err := json.Marshal(postBody)
	if err != nil{
		log.Println(err)
	}

	// Send POST request to ECE
	client := &http.Client{}
	URI := HTTP_PROTOCOL + eceCoordinatorHost + ":" + ECE_PORT + TOKEN_ENDPOINT
	request, _ := http.NewRequest("POST", URI, bytes.NewBuffer(jsonPayload))
	request.Header.Add("Content-Type", "application/json")
	request.SetBasicAuth(eceUser, ecePassword)
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	roleToken := RoleTokenResponse{}
	err = json.Unmarshal(body, &roleToken)
	if err != nil {
		return nil, err
	}
	return &roleToken, nil
}