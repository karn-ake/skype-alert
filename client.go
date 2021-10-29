package main

import (
	"fmt"

	"./repository"
)

const urlblp = "http://192.168.142.51:8080/api/blp"

func main() {
	data, _ := repository.FetchData(urlblp)

	fmt.Println(data)
}
