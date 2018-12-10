package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {

	defer func() {

		if r := recover(); r != nil {
			// utillog.Instance().Fatal(r)
		}
	}()

	// panic(55)

	fmt.Println("启动服务器.....")

	// var wg sync.WaitGroup
	// for i := 0; i < 500; i++ {
	// 	wg.Add(1)
	// 	go httpGet1()
	// }
	// wg.Wait()

	ch := make(chan string)

	FileNum := 300
	// 每个 goroutine 处理一个文件的下载

	for i := 0; i < FileNum; i++ {
		go func(fileID int) {
			ch <- Download(i)
		}(i)
	}

	// 等待每个文件下载的完成，并检查超时
	timeout := time.After(900 * time.Second)
	for idx := 0; idx < FileNum; idx++ {
		select {
		case res := <-ch:
			nt := time.Now().Format("2006-01-02 15:04:05")
			fmt.Printf("[%s]Finish download %s\n", nt, res)
		case <-timeout:
			fmt.Println("Timeout...")
			break
		}
	}
}

// 实现单个文件的下载
func Download(fileID int) string {
	nt := time.Now().Format("2006-01-02 15:04:05")

	id := GetGID()
	fmt.Printf("[%s]To download %s\n", nt, id)

	resp1, err := http.Get("http://192.168.1.169:8080/CheckUpdateInfoJson?name=bo-android&version=1")

	if err != nil {

		// handle error
		fmt.Printf("111100 %s ", err)
	}

	defer resp1.Body.Close()

	body, err := ioutil.ReadAll(resp1.Body)

	if err != nil {

		// handle error
		fmt.Println("1111022")
	}

	fmt.Println(string(body))

	// url := "http://13.250.51.60:8080/downloadFile?name=bo-android"
	url := "http://192.168.1.169:8080/static/versions/bo-android/4.apk"
	fpath := fmt.Sprintf("./aaa/%s", id)
	newFile, err := os.Create(fpath)
	if err != nil {
		fmt.Println(err.Error())
		return "process failed for "
	}
	defer newFile.Close()

	client := http.Client{Timeout: 900 * time.Second}
	resp, err := client.Get(url)
	defer resp.Body.Close()

	_, err = io.Copy(newFile, resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return "fileID"
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func httpGet() {

	resp, err := http.Get("https://192.168.1.169:8080/CheckUpdateInfoJson?name=bo-android&version=1")

	if err != nil {

		// handle error

	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		// handle error

	}

	fmt.Println(string(body))

}
