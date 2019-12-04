package main

import (
    "fmt"
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

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
            "password=%s dbname=%s sslmode=disable",
            os.Getenv("pghost"),
            os.Getenv("pgport"),
            os.Getenv("pguser"),
            os.Getenv("pgpwd"),
            "helm_demo")

        fmt.Fprintf(w, "Ip %s processing request\n", GetLocalIP())
        fmt.Fprintf(w, "psqlInfo %s\n", psqlInfo)
    })

    http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Path[len("/employee/"):]
        psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
            "password=%s dbname=%s sslmode=disable",
            os.Getenv("pghost"),
            os.Getenv("pgport"),
            os.Getenv("pguser"),
            os.Getenv("pgpwd"),
            "helm_demo")

        fmt.Fprintf(w, "Ip %s processing request\n", GetLocalIP())
        fmt.Fprintf(w, "employee %s\n", name)
        fmt.Fprintf(w, "psqlInfo %s\n", psqlInfo)
    })

    http.ListenAndServe(":8080", nil)
}
