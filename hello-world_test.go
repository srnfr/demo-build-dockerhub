package main

import (
    "testing"
    "github.com/valyala/fastjson"
)

func TestGetVer(t *testing.T) {
    version := getVer()
    if version < 10 {
        t.Fatalf(`getVer("") = %v`, version)
    }
}

func TestGetIP(t *testing.T) {
    result := getIP()
    test := fastjson.Validate(result)
    if test != nil {
        t.Fatalf(`getIP("") = %v`, test)
    }
}