package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func bar() string {
	return "Bar"
}

func main() {
	fmt.Println(bar())
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Errorf("[ERR] Failed to get\n")
		fmt.Errorf("[Err] %v\n", err.Error())
	} else {
		respBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Errorf("[ERR] Failed to get\n")
			fmt.Errorf("[Err] %v\n", err.Error())

		} else {
			fmt.Println(string(respBody))
		}
	}
}
