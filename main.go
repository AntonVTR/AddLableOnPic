package main

import (
	"flag"
	"fmt"

	P "github.com/AntonVTR/AddLableOnPic/Library"
)

func main() {
	url := flag.String("url", "", "URL to the file with result")
	flag.Parse()
	fmt.Println("url is", *url)

	P.ParseU(*url)
	//P.ParseU("https://raw.githubusercontent.com/AntonVTR/AddLableOnPic/master/results/protocol2019.htm")

}
