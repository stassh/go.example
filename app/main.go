package main

import (
    // "fmt"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func homePage(w http.ResponseWriter, r *http.Request){

  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(`{"message": "hello world"}`))

  // fmt.Fprintf(w, "Welcome to the HomePage!")
  // fmt.Println("Endpoint Hit: homePage")
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(`{"message": "welcome"}`))
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(`{"message": "welcome ` + ps.ByName("name") + `"}`))
  // fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func handleRequests() {
    // http.HandleFunc("/", homePage)

    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":30080", router))
}

func main() {
    handleRequests()
}