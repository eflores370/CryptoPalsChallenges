package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func parseCookie(cookies string) {

	cookieArray := strings.Split(cookies, string('&'))
	object := map[string]string{}
	for _,v := range cookieArray {
		keyValuePair := strings.Split(v, string('='))
		object[keyValuePair[0]] = keyValuePair [1]
	}
	jsonObject,_ := json.Marshal(object)
	fmt.Println(string(jsonObject))
}

func main(){
	cookies := "foo=bar&baz=qux&zap=zazzle"

	parseCookie(cookies)
}