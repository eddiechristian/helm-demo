package main

import (
    "fmt"
    "net/http"
    "net"
    "database/sql"
    _ "github.com/lib/pq"
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
        fmt.Fprintf(w, "Ip %s processing request\n", GetLocalIP())
    })

    http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Path[len("/employee/"):]
        fmt.Fprintf(w, "Ip %s processing request\n", GetLocalIP())
        fmt.Fprintf(w, "employee %s\n", name)
    })

    http.ListenAndServe(":8080", nil)
}
