package main

import (
    "database/sql"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "time"
)

var dataService DataService

type IndexPageModel struct {
    AllShop    []ShopInfo
    AllAccount []AccountInfo
}

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

var routes = []Route{
    Route{
        "Index",
        "Get",
        "/",
        Index,
    },
    Route{
        "Test",
        "Get",
        "/test",
        Test,
    },
    Route{
        "AllAccount",
        "Get",
        "/all-account",
        GetAllAccount,
    },
}

func main() {
    log.Println("creditease account map starting...")

    startHttpServer := func(addr string) {
        router := NewRouter()
        log.Fatal(http.ListenAndServe(addr, router))
        log.Println("server stoped !!!")
    }

    // init dataService:
    dbFile := "./creditease.db"
    db, err := sql.Open("sqlite3", dbFile)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    dataService = &DataServiceImplSqlite3{dbFile, db}

    if len(os.Args) > 1 {
        if os.Args[1] == "--import-db" {
            ImportDb(os.Args[2:], dataService)
        } else if os.Args[1] == "--import-js" {
            GenerateJs(os.Args[2:])
        } else if os.Args[1] == "--port" {
            serverAddr := ":9000"
            if len(os.Args) > 2 {
                port := os.Args[2]
                serverAddr = ":" + port
            }
            startHttpServer(serverAddr)
        } else {
            log.Println("Invalid parameter: ", os.Args[1])
        }
    } else {
        startHttpServer(":9000")
    }
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    s := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))
    router.PathPrefix("/assets/").Handler(s)
    for _, route := range routes {
        var handler http.Handler

        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}

func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t%s\t%s\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
