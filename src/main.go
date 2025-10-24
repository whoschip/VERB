package main

import (
	"VERB/module/check"
	"VERB/module/getbio"
	"VERB/module/listdownloader"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Downloading word list")
	listdownloader.Download("https://raw.githubusercontent.com/whoschip/wordlist/main/data/all.txt")

	data, err := ioutil.ReadFile("wordlist.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	funny := getbio.GetbioWithCheck(7513451449)

	if funny == "banned" {
		fmt.Print("yo roblox good job vro")
	} else {
		matches := check.Check(funny, string(data))

		fmt.Println("Found words:", matches)
	}

}
