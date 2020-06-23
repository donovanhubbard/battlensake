package main

import (
    "log"
    "os"
    "net/http"
    "fmt"
)

func Index( w http.ResponseWriter, r *http.Request){
    log.Print("Hit the index")
    fmt.Fprintf(w, "Hi there")
}

func main(){

    http.HandleFunc("/",Index)

    port := os.Getenv("PORT")
    if port == ""{
        port = "80"
    }

    log.Printf("Starting program on port %s\n",port)

    http.ListenAndServe(":"+port,  nil)

}
