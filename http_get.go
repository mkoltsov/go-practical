package main

import (
"io/ioutil"
"fmt"
"net/http")

func main() {
	resp,_:=http.Get("http://javabean.ru")
	body,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}