package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetImg(url string) (n int64, err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = "img/" + path[len(path)-1]
	}
	fmt.Println("path :", path)
	fmt.Println("name :", name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	return
}
