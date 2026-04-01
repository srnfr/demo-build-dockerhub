package main

// Imports (non nécessaires)
import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func getIP() string {
    // L'API est payant en HTTPS mais tjrs Free en HTTP (??)
    response, err := http.Get("http://ip-api.com/json/?fields=query")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    return string(responseData)

}

func getVer() int {
    version:= 123
    return version
}

func main() {

    fmt.Println("Return from ip-api.com: ")
    fmt.Println(string(getIP()))

    version := getVer()
    fmt.Println(`(new) hello world version`,version)
}
