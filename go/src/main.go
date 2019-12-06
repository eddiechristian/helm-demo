package main

import (
    "fmt"
    "database/sql"
    "net/http"
    "net"
    _ "github.com/lib/pq"
    "os"
)

func GetLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

//example https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
            "password=%s dbname=%s sslmode=disable",
            os.Getenv("pghost"),
            os.Getenv("pgport"),
            os.Getenv("pguser"),
            os.Getenv("pgpwd"),
            "helm_demo")
        
        fmt.Fprintf(w, "Ip %s processing request\n", GetLocalIP())
        db, err := sql.Open("postgres", psqlInfo)
        if err != nil {
            panic(err)
        } else {
            rows, err := db.Query("SELECT username, email FROM account")
            if err != nil {
                panic(err)
            } else {
                fmt.Fprintf(w, "username | email\n")
                for rows.Next() {
                    var username string
                    var email string
                    err = rows.Scan(&username, &email)
                    fmt.Fprintf(w, "%8v | %8v\n", username, email)
                }
            }
        }
        
    })

    http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Path[len("/employee/"):]
        psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
            "password=%s dbname=%s sslmode=disable",
            os.Getenv("pghost"),
            os.Getenv("pgport"),
            os.Getenv("pguser"),
            os.Getenv("pgpwd"),
            "helm_demo")

        fmt.Fprintf(w, "Ip %s processing request\n", GetLocalIP())
        fmt.Fprintf(w, "employee %s\n", name)
        db, err := sql.Open("postgres", psqlInfo)
        if err != nil {
            panic(err)
        } else {
            query := fmt.Sprintf("SELECT username, email FROM account where username like '%%%s%%'", name)
            fmt.Fprintf(w, "query %s\n", query)
            rows, err := db.Query(query)
            if err != nil {
                panic(err)
            } else {
                fmt.Fprintf(w, "username | email\n")
                for rows.Next() {
                    var username string
                    var email string
                    err = rows.Scan(&username, &email)
                    fmt.Fprintf(w, "%8v | %8v\n", username, email)
                }
            }
        }
    })

    http.ListenAndServe(":8080", nil)
}
