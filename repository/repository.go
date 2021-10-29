package repository

import (
	"encoding/json"
	"net/http"

	"../entity"
)

func FetchData(url string) (*entity.Customer, error) {
	var getdata = &http.Client{}
	data, _ := getdata.Get(url)
	var cust entity.Customer
	err := json.NewDecoder(data.Body).Decode(&cust)
	if err != nil {
		return &cust, err
	}
	return &cust, nil
}
