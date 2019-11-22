package main

import (
    "fmt"
    "net/http"
    "net"
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

    http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Path[len("/greet/"):]
        fmt.Fprintf(w, "Hello %s\n", name)
    })

    http.ListenAndServe(":8080", nil)
}
