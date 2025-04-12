package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var URL = "https://httpbin.org/headers"

func main() {
	client := http.Client{}

	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("X-Client-Name", "Sardor")

	r, _ := client.Do(req)

	defer r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	fmt.Println(string(body))
	// "ADD:4,6"
	// strings.Split("", ":") -> ["ADD", "4,6"]
	// 
	// strings.split("4,6", ",") -> ["4", "6"]
	// a,_ = strconv.ParseInt("4") -> 4
	// b,_ = strconv.ParseInt("6") -> 6


	// COMMAND:DATA
	// "ECHO:ILFAT"
	// "TIME"
	// "EXIT"
}

/*

	ADD -> math operation
	5,5 -> values (a, b)


	"SUB:4,6"

	"OPERATION" "VALUES"
	VALUES -> "value,value"
	VALUE VALUE

	"23" -> 23
	"5" -> 5
	"A5" -> error


*/
