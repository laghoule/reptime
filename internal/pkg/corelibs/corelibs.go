/*
Here will be a description of the tools
*/

package corelibs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetBody(target string) string {
	fmt.Printf("My target: %s\n", target)

	res, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
