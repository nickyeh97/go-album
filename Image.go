package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var folderPath = "img/"

func GetImg(url string) (err error) {
	path := strings.Split(url, "/")
	var name string
	if len(path) > 1 {
		name = folderPath + path[len(path)-1]
	}
	fmt.Println("path :", path)
	fmt.Println("name :", name)

	os.MkdirAll("img", os.ModePerm)

	// 創建寫入檔案串流的io.Writer
	out, err := os.Create(name)
	if err != nil {
		return err
	}
	defer out.Close()
	fmt.Println("out: ", out)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("resp: ", resp)
	// base64 data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	n, err := io.Copy(out, bytes.NewReader(data))
	if err != nil {
		return err
	}
	fmt.Println("n: ", n)

	return
}
