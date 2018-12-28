package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte{'h','e','l','l','o',' ','w','o','r','l','d'})
	})

	http.ListenAndServe(":9002", nil) //设置监听的端口
}
