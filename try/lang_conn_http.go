package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const HttpAddr = "http://127.0.0.1/aa"

func main() {
	go serviceHttp()

	clientHttp()
}

func clientHttp() {
	sw := sync.WaitGroup{}
	sw.Add(2)

	go func() {
		defer sw.Done()
		clientCallApi(HttpAddr)
	}()
	go func() {
		defer sw.Done()
		clientCallApi2(HttpAddr)
	}()
	sw.Wait()
}

func clientCallApi(path string) {
	res, err := http.Get(path)
	if err != nil {
		fmt.Printf("err : %s\n", err)
		return
	}
	defer func() {
		_ = res.Body.Close()
	}()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("接受服务端1：", string(data))
}

func clientCallApi2(path string) {
	fmt.Println("path：", path)
	client := &http.Client{}
	request, err := http.NewRequest("GET", path, bytes.NewBuffer([]byte("abcssss")))
	if err != nil {
		fmt.Printf("err : %s\n", err)
		return
	}
	res, err := client.Do(request)
	defer func() {
		_ = res.Body.Close()
	}()
	reader := bufio.NewReader(res.Body)
	for {
		res, err := reader.ReadString(' ')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("NewReader err : %s\n", err)
			return
		}
		fmt.Println("接受服务端2：", string(res))
	}
}

func serviceHttp() {
	fmt.Println("服务端")
	http.HandleFunc("/aa", serviceApi)
	_ = http.ListenAndServe(":80", nil)

}

func serviceApi(w http.ResponseWriter, r *http.Request) {
	i := 1
	for {
		data := fmt.Sprintf("hello word %d ", i)
		_, _ = fmt.Fprintf(w, data)
		i++
		if i > 10 {
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}
