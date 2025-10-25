package main

import (
	"fmt"
	"os"

	"github.com/whoschip/WordsCheck/module/check"
	"github.com/whoschip/WordsCheck/module/listdownloader"
)

func Check(bio string, wordslist string) []string {
	fmt.Println("Downloading word list")
	listdownloader.Download("https://raw.githubusercontent.com/whoschip/wordlist/main/data/all.txt")

	data, err := os.ReadFile("wordlist.txt")
	if err != nil {
		fmt.Println("err:", err)
		return nil
	}

	matches := check.Check(bio, string(data))
	return matches
}
