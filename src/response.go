package main

import (
    "package"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello"))
}
 
func main() {
    http.HandleFunc("/hello", SayHello)
    http.ListenAndServe(":8001", nil)
 
}

func ajax(data NewBaseJsonBean)  {
    
}