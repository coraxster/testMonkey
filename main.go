package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"./config"
	"encoding/json"
	"net/http"
	"time"
)

func main(){
	var conf config.Config
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	e = json.Unmarshal(file, &conf)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	for _, ep := range conf.Endpoints {
		http.HandleFunc(ep.Uri, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(ep.Status)
			w.Write(ep.Response)

			filepath := "./requests/" + time.Now().String()
			os.MkdirAll("./requests/", os.ModePerm)
			var body []byte
			r.Body.Read(body)
			fo, _ := os.Create(filepath)
			fo.Write(body)
		})
	}
	if err := http.ListenAndServe(conf.Bind, nil); err != nil {
		panic(err)
	}
}
