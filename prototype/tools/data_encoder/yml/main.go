package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ritsource/episteme/prototype/data/models"
	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 2 {
		panic("please pass an arguement with source yml file path")
	}
	ymlfp := os.Args[1]
	fmt.Printf("ymlfp = %+v\n", ymlfp)

	data, err := ioutil.ReadFile(ymlfp)
	if err != nil {
		fmt.Printf("unable to read file %v\n", ymlfp)
		panic(err)
	}

	posts := models.Posts{}

	err = yaml.Unmarshal(data, &posts)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", posts)
}
