package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

// Engineer - Our struct for all Engineers
type Engineer struct {
    Id      string    `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

var Engineers []Engineer

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllEngineers(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllEngineers")
    json.NewEncoder(w).Encode(Engineers)
}

func returnSingleEngineer(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, engineer := range Engineers {
        if engineer.Id == key {
            json.NewEncoder(w).Encode(engineer)
        }
    }
}


func createNewEngineer(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Engineer struct
    // append this to our Engineers array.
    reqBody, _ := ioutil.ReadAll(r.Body)
    var engineer Engineer
    json.Unmarshal(reqBody, &engineer)
    // update our global Engineer array to include
    // our new Engineer
    Engineers = append(Engineers, engineer)

    json.NewEncoder(w).Encode(engineer)
}

func deleteEngineer(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, engineer := range Engineers {
        if engineer.Id == id {
            Engineers = append(Engineers[:index], Engineers[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/engineers", returnAllEngineers)
    myRouter.HandleFunc("/engineer", createNewEngineer).Methods("POST")
    myRouter.HandleFunc("/engineer/{id}", deleteEngineer).Methods("DELETE")
    myRouter.HandleFunc("/engineer/{id}", returnSingleEngineer)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    Engineers = []Engineer{
        Engineer{Id: "1", Title: "Mehmet", Desc: "Engineer Description 1", Content: "Engineer Content 1"},
        Engineer{Id: "2", Title: "Eray ", Desc: "Engineer Description 2", Content: "Engineer Content 2"},
        Engineer{Id: "3", Title: "Kaan ", Desc: "Engineer Description 3", Content: "Engineer Content 3"},
        Engineer{Id: "4", Title: "Melih 2", Desc: "Engineer Description 4", Content: "Engineer Content 4"},
    }
    handleRequests()
}