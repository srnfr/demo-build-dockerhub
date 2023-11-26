package main

// Imports (non nécessaires)
import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {

    response, err := http.Get("http://ip-api.com/json/?fields=query")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))


    fmt.Println("hello world v82")
}
