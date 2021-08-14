package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Message struct {
	Name string
	Body string
	Time int64
}

var json_example = `{
    "args": {},
    "headers": {
        "x-forwarded-proto": "https",
        "x-forwarded-port": "443",
        "host": "postman-echo.com",
        "x-amzn-trace-id": "Root=1-6117a983-446ddff96dedbad80be30d58",
        "user-agent": "PostmanRuntime/7.28.2",
        "accept": "*/*",
        "postman-token": "a402b2f1-2686-459d-a16e-e7a50dd986af",
        "accept-encoding": "gzip, deflate, br"
    },
    "url": "https://postman-echo.com/get"
}`

func main() {
	fmt.Println(json_example)
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	fmt.Printf("%s\n", b)
	var m2 Message
	json.Unmarshal(b, &m2)
	fmt.Println(reflect.DeepEqual(m, m2))
}
