package main

import (
    "testing"
)

func TestGetVer(t *testing.T) {
    version := getVer()
    if version < 10 {
        t.Fatalf(`getVer("") = %v`, version)
    }
}