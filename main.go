package main

import (
    "errors"
    "fmt"
    "log"
    "net/http"
)

type Person struct {
    Name    string
    Age     int
}

func personCreate(w http.ResponseWriter, r *http.Request) {
    var p Person

    err := decodeJSONBody(w, r, &p)
    if err != nil {
        var mr *malformedRequest

        if errors.As(err, &mr) {
            // If the error is of type that we created, then we just display it to user
            http.Error(w, mr.msg, mr.status)
        } else {
            // Otherwse, this is an internal server error
            log.Println(err.Error())
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        }

        return
    }

    fmt.Fprintf(w, "Person: %+v", p)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/person/create", personCreate)

    if err := http.ListenAndServe(":4000", mux); err != nil {
        log.Fatal(err)
    }
}
