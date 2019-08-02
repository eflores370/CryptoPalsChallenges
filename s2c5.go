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

func profile_for(email string) (cookie string) {
	cookieMap := map[string]string{"email": email, "uid": "10", "role":"user"}
	
	//fmt.Println(cookieMap)
	for key,v := range cookieMap {
		cookie += key+"="+v+"&"
	}

	return cookie[:len(cookie)-1]
}

func main(){
	cookies := "foo=bar&baz=qux&zap=zazzle"
	parseCookie(cookies)
	fmt.Println(profile_for("foo@bar.com"))
}