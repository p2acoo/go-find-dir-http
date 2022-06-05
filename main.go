package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var out *string
var url *string
var number int

func main() {
	getFlag()
}

func getFlag() {
	help := flag.Bool("help", false, "show this help")
	out := flag.String("file", "", "file to read")
	url := flag.String("url", "", "url to test")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	} else if *out != "" && *url != "" {
		fmt.Println("Start")
		readFile(*out, *url)
	}
}

func readFile(files string, urls string) {
	fmt.Println("Read file")
	file, err := os.Open(files)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number++
		get, _ := http.Get(urls + scanner.Text())
		if get.StatusCode == 404 || get.StatusCode == 401 {

		} else {
			fmt.Println("OK " + urls + scanner.Text() + " " + get.Status)
			go readFile(files, urls+scanner.Text()+"/")
		}
		fmt.Println(number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
