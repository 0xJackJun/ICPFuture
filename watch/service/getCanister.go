package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"watch/model"
)

// GetCanister Get all monitored Canister IDs in node
func GetCanister()  []model.User {
	// local command
	args := []string{"canister", "call", "r7inp-6aaaa-aaaaa-aaabq-cai" ,"get"}
	cmd := exec.Command("dfx", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if  err != nil {
		log.Fatal(err)
	}
	return Regular(string(opBytes))
}

func Regular(op string)  []model.User {
	reg1 := regexp.MustCompile(`= "(.*?)";`)
	if reg1 == nil {
		fmt.Println("err = ")
		return nil
	}

	result := reg1.FindAllStringSubmatch(op, -1)
	users := make([]model.User,0,2)
	for i := 0; i<len(result);i +=3 {
		user := model.User{}
		user.Threshold = result[i][1]
		user.CanisterID = result[i+1][1]
		user.Email = result[i+2][1]
		users = append(users,user)
	}
	return users
}

