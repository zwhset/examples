package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() {
			fs, _ := ioutil.ReadDir(filepath.Join(".", file.Name()))
			for _, f := range fs {
				fmt.Printf("\t %s\n", f.Name())
			}
		}
	}

	content, err := ioutil.ReadFile("./k8s/client.go")
}
