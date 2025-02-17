package main

import (
	"fmt"
	"github.com/IPQualityScore/GoEmailDBReader/pkg/emaillookup"
	"time"
)

func main() {
	input := ""
	lookup := emaillookup.EmailLookup{Path: "./tree/"}
	for true {
		fmt.Println("Enter email to search:")
		fmt.Scanln(&input)
		start := time.Now()
		res := lookup.LookupEmail(input)
		if res != nil {
			fmt.Println("found data:", res)
			for _, v := range res.Data {
				fmt.Println(v.ToString())
			}
			fraud := res.FraudScore()
			if fraud != nil {
				fmt.Println("parsed fraudscore:", fraud.FraudScore)
			} else {
				fmt.Println("couldn't find fraudscore")
			}
		} else {
			fmt.Println("not found")
		}
		fmt.Println("Time taken:", time.Since(start))
	}
}
