package main

import (
    "testing"
    "github.com/valyala/fastjson"
)

func TestGetVer(t *testing.T) {
    // Verifie que la version est supérieure à 10
    version := getVer()
    if version < 10 {
        t.Fatalf(`getVer("") = %v`, version)
    }
}

func TestGetIPValidJSON(t *testing.T) {
    // Verifie que le JSON est valide
    result := getIP()
    test := fastjson.Validate(result)
    if test != nil {
        t.Fatalf(`getIP("") = %v`, test)
    }
}

func TestGetIPQueryKeyExists(t *testing.T) {
    //Verifie qu'il existe un champ query dans le JSON
    result := getIP()
    want := true
    test := fastjson.Exists([]byte (result), "query")
    if test != want {
        t.Fatalf(`getIP("") = %v`, test)
    }
}