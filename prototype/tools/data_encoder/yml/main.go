package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ritsource/episteme/prototype/data/models"
	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 2 {
		panic("no input file provided")
	}
	if len(os.Args) < 3 {
		panic("no output file provided")
	}
	ifp := os.Args[1]
	ofp := os.Args[2]
	fmt.Printf("ymlfp = %+v\n", ifp)
	fmt.Printf("outfp = %+v\n", ofp)

	if ifp == "" {
		fmt.Printf("invalid input file path %+v\n", ifp)
	}
	if ofp == "" {
		fmt.Printf("invalid output file path %+v\n", ofp)
	}

	data, err := ioutil.ReadFile(ifp)
	if err != nil {
		fmt.Printf("unable to read file %v\n", ofp)
		panic(err)
	}

	posts := models.Posts{}
	err = yaml.Unmarshal(data, &posts)
	if err != nil {
		fmt.Printf("unable to decode %v data\n", ifp)
		panic(err)
	}

	err = posts.SaveToFS(ofp)
	if err != nil {
		fmt.Printf("unable to save message data file to %+v\n", ofp)
		panic(err)
	}

}
