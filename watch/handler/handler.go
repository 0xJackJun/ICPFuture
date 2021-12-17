package handler

import (
	"fmt"
	"github.com/golang/groupcache/consistenthash"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"regexp"

	"github.com/gin-gonic/gin"

	"watch/model"
)

func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func SubmitInfo(c *gin.Context) {
	canisterID := c.Query("canisterID")
	threshold := c.Query("threshold")
	email := c.Query("email")
	if !VerifyEmail(email) {
		fmt.Println("Incorrect email format")
	}
	SendMonitInfo(canisterID,threshold,email)
	SendResponse(c, nil)
}
func SendMonitInfo(canisterID,threshold,email string)  {
	node := LoadBalancing(canisterID)
	Splicing := "'(record {canister_id=" + canisterID + ";" +
		"email="+ email+";"+"threshold="+threshold+"})'"
	args := []string{"canister", "call", node, "create", Splicing}
	// dfx canister call ryjl3-tyaaa-aaaaa-aaaba-cai create
	// '(record {canister_id="ryjl3-tyaaa-aaaaa-aaaba-cai";
	// email="wzzytu@163.com";threshold="1234567890"})'
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
	fmt.Println("Successful upload of monitored Canister：",string(opBytes))
}

func SendResponse(c *gin.Context, err error) {
	code, message := decodeErr(err)
	c.JSON(http.StatusOK, model.Response{
		Result:    err == nil,
		Code:      code,
		Message:   message,
	})
}

func decodeErr(err error) (string, string) {
	if err == nil {
		return "0", "SUCCESS"
	} else {
		return "-1", "FAIL" + err.Error()
	}
}

func AddNode(c *gin.Context) {
	node := c.Query("canisterID")
	AddNodes(node)
	SendResponse(c, nil)
}

var Nodes = make([]string,0,2)

// AddNode Add Canister, used to store monitored Canister information
func AddNodes(canisterID string) {
	Nodes = append(Nodes, canisterID)
}


// LoadBalancing Distribute the monitored containers
// to different Canisters for storage to achieve the
// goal of load balancing
func LoadBalancing (moniCanister string) string {
	hash := consistenthash.New(3, nil)
	for _, node := range Nodes {
		if len(node) == 0 {
			continue
		}
		hash.Add(
			node,
		)
	}
	node := hash.Get(moniCanister)
	return node
}