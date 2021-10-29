package main

import (
	"fmt"
	"log"

	"./repository"
	"github.com/andrewdruzhinin/go-skype/skype"
)

const urlblp = "http://192.168.142.51:8080/api/blp"

func main() {
	data, _ := repository.FetchData(urlblp)
	fmt.Println(data)

	s := skype.NewClient("karn-ake.r@n2nconnect.com", "N2Nconnect123!@#")
	//s := skype.NewClient("live:.cid.e96d2c346fc1dcd3", "N2Nconnect123!@#")
	_, err := s.Authorization.Authorize()
	if err != nil {
		log.Fatal(err)
		return
	}

	//	fmt.Println(a)

	_, err1 := s.Messege.Send("19:509e09dc857444e88e1b0b3764f7b818@thread.skype", "text", "Test")
	//_, err1 := s.Messege.Send("19:509e09dc857444e88e1b0b3764f7b818@thread.skype", "text", "Test")
	if err != nil {
		fmt.Println(err1)
		return
	}
	//	fmt.Println(*b)

}
