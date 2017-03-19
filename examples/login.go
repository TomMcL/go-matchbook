package main

import (
	"github.com/tommcl/matchbook"
	"log"
	"os"
	"encoding/json"
	"fmt"
)

var confFile = "YOUR_FILE_PATH_HERE"

func main() {
	file, err := os.Open(confFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	login := new(matchbook.Config)
	err = dec.Decode(&login)
	if err != nil {
		log.Fatal(err)
	}

	sess, err := login.NewSession()
	if err != nil{
		log.Fatal(err)
	}

	resp, err := sess.Login()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response is %s", resp)
}
