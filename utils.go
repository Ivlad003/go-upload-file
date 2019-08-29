package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func saveFileFromURL(url string, w http.ResponseWriter) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	cropFile(w, tempFile)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	var s []DataImage
	r.ParseMultipartForm(32 << 20)
	fhs := r.MultipartForm.File["myfiles"]
	for _, fh := range fhs {
		file, err := fh.Open()

		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		s = append(s, cropFile(w, tempFile))
	}

	sendJson(w, s)
}

func cropFile(w http.ResponseWriter, f *os.File) DataImage {
	fmt.Println("cImg dimension:", f.Name())
	src, err := imaging.Open(f.Name())
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dst := imaging.Resize(src, 100, 100, imaging.Lanczos)
	dir, file := filepath.Split(f.Name())
	err = imaging.Save(dst, filepath.Join(dir, "100x100"+file))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	dataImage := DataImage{file, "100x100" + file}

	return dataImage
}

func sendJson(w http.ResponseWriter, dataImage []DataImage) {
	js, err := json.Marshal(dataImage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
