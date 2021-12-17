package service

import (
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetCycles(node, canisterId string) int {
	//node := LoadBalancing(canisterId)
	Splicing := "(record{canister_id=principal " + "\"" + canisterId + "\"" + "})"
	args := []string{"canister", "call", node, "getCycle", Splicing}
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
	var cycleBalance int
	if len(opBytes) != 0 {
		res := string(opBytes)
		a := strings.Index(res, "(")
		b := strings.Index(res, ":")
		rowStr := res[a+1:b]
		split := strings.Split(rowStr, "_")
		joinStr := strings.Join(split, "")
		g := strings.TrimSpace(joinStr)
		len := len(g)
		cyclesStr := g[:len-1]
		cycleBalance, _ = strconv.Atoi(cyclesStr)

	}
	return cycleBalance
}
