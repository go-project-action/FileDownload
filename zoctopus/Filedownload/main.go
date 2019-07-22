package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

//声明要下载的文件地址
var URL = "https://dl.google.com/go/go1.10.3.darwin-amd64.pkg"

func main() {
	//解析文件地址
	uri, err := url.ParseRequestURI(URL)
	if err != nil {
		fmt.Println("URL Error", err.Error())
	}

	filename := path.Base(uri.Path) //读取文件地址的文件名称
	log.Println("[*]Filename" + filename)

	/*
		创建一个httpClient
	*/
	client := http.DefaultClient
	client.Timeout = time.Second * 30 //设置超时时间30s
	resp, resperr := client.Get(URL)
	if resperr != nil {
		fmt.Println("resp Error", resperr.Error())
	}

	raw := resp.Body
	defer raw.Close()
	reader := bufio.NewReaderSize(raw, 20*1024)

	file, cferr := os.Create(filename)
	if cferr != nil {
		fmt.Println("Create File Error", cferr.Error())
	}
	writer := bufio.NewWriter(file)

	buff := make([]byte, 20*1024)
	written := 0

	/*
		开协程（重点）
	*/
	var downloaderr error
	go func() {
		for {
			nr, err := reader.Read(buff)
			if nr > 0 {
				nw, ew := writer.Write(buff[0:nr])
				if nw > 0 {
					written += nw
				}
				if ew != nil {
					downloaderr = ew
					break
				}
				if nr != nw {
					downloaderr = io.ErrShortWrite
					break
				}
			}
			if err != nil {
				if err != io.EOF {
					downloaderr = err
				}
			}
		}
		if err != nil {
			fmt.Println("erro", err.Error())
		}
	}()

	spacetime := time.Second * 1

	ticker := time.NewTicker(spacetime) //设置定时器

	lastwritelength := 0

	stop := false //设置启动状态

	/*
		堵塞住程序，然后定时器按照间隔时间来计算速度
	*/

	for {
		select {
		case <-ticker.C:
			speed := written - lastwritelength //注意，如果间隔时间不为1，需要做除法算下载速度
			fmt.Printf("[*] Speed %s / %s \n", bytesToSize(speed), spacetime.String())
			if speed == 0 { //说明文件下载完了
				ticker.Stop()
				stop = true
				break
			}
			lastwritelength = written
		}
		if stop {
			break
		}
	}

}

func bytesToSize(length int) string {
	var k = 1024 // or 1024
	var sizes = []string{"Bytes", "KB", "MB", "GB", "TB"}
	if length == 0 {
		return "0 Bytes"
	}
	i := math.Floor(math.Log(float64(length)) / math.Log(float64(k)))
	r := float64(length) / math.Pow(float64(k), i)
	return strconv.FormatFloat(r, 'f', 3, 64) + " " + sizes[int(i)]
}
