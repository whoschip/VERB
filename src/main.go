package main

import (
	"WordsCheck-VERB/module/check"
	"WordsCheck-VERB/module/getbio"
	"WordsCheck-VERB/module/listdownloader"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Downloading word list")
	listdownloader.Download("https://raw.githubusercontent.com/whoschip/wordlist/main/data/all.txt") // my version

	data, err := ioutil.ReadFile("wordlist.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	var lolid int

	fmt.Println("Roblox user id to scan :")

	fmt.Scanln(&lolid)

	funny := getbio.GetbioWithCheck(lolid) // user id i think

	if funny == "banned" {
		fmt.Print("yo roblox good job vro")
	} else {
		matches := check.Check(funny, string(data))

		fmt.Println("Found words:", matches)
	}

}
