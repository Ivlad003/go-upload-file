package main

import (
	"encoding/json"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
)

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, filepath.Join("temp-images", path.Base(r.URL.Path)))
	case "PUT":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		mapUrl := map[string]string{}
		err = json.Unmarshal(body, &mapUrl)
		if err != nil {
			panic(err)
		}
		saveFileFromURL(mapUrl["url"], w)
	case "POST":
		uploadFile(w, r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func setupRoutes() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}

// go run main.go models.go utils.go

// curl -X POST -F myfiles=@test.jpg http://localhost:8080
// {"Name":"upload-351526307.png","CropName":"100x100upload-351526307.png"}

// curl -H 'Content-Type: application/json' -X PUT -d "{\"url\":\"http://i.imgur.com/m1UIjW1.jpg\"}" http://localhost:8080/
// curl: (6) Could not resolve host: application
// {"Name":"upload-174286797.png","CropName":"100x100upload-174286797.png"}
