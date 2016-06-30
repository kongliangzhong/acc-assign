package main

import (
    "net/http"
    "html/template"
    "encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("assets/indexPage.html")
    err := t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func Test(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("assets/testPage.html")
    err := t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func GetAllAccount(w http.ResponseWriter, r *http.Request) {
    allAccount := dataService.GetAllAccount()
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(allAccount); err != nil {
        panic(err)
    }
}

func GetAllShop(w http.ResponseWriter, r *http.Request) {

}
