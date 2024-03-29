package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("soffice.exe", "--pt", os.Getenv("POS_PRINTER"), os.Args[1])
	_, err := cmd.Output()
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	os.Exit(0)
}
