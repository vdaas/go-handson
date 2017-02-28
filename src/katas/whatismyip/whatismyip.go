package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// str, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(str))

	result := make(map[string]string)
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("My IP address is:", result["origin"])
}
