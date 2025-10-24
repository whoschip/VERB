package main

import (
	"fmt"
	"goon-detector/module/check"
	"goon-detector/module/getbio"
	"goon-detector/module/listdownloader"
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

	funny := getbio.GetbioWithCheck(8386732330)

	if funny == "banned" {
		fmt.Print("yo roblox good job vro")
	} else {
		matches := check.Check(funny, string(data))

		fmt.Println("Found words:", matches)
	}

}
