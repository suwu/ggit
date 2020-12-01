package main

import (
	"log"
	"strings"
	"os"
	"os/exec"
)

const github_host = "https://github.com"
const replace_host = "https://github.com.cnpmjs.org"

func main() {
	if len(os.Args) >= 3 {
		if os.Args[1] == "clone" {
			url := os.Args[2]

			log.Println("old url:", url)
			url = strings.Replace(url, github_host, replace_host, 1)
			log.Println("new url:", url)
			os.Args[2] = url
		}
	}

	cmd := exec.Command("git", os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
