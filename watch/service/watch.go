package service

import (
	"fmt"
	"strconv"
	"watch/handler"
)

func Watch() {
	for _, node := range handler.Nodes {
		items := GetCanister()
		for _, item := range items {
			if item.CanisterID == "" {
				continue
			}
			cycles := GetCycles(node,item.CanisterID)
			threshold,err := strconv.Atoi(item.Threshold)
			if err != nil {
				fmt.Println("The obtained threshold cannot be converted to int type")
			}
			if cycles < threshold {
				SendMail(item.Email)
			}
		}
	}
}

