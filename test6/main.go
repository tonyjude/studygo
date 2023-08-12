package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	chSem      = make(chan int, 3)
	downloadWG sync.WaitGroup
)

func DownloadAsync(url string) {
	downloadWG.Add(1)
	go func() {
		chSem <- 123
		Download(url)
		<-chSem
		downloadWG.Done()
	}()
}

func Download(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("net error=", err)
	}
	defer resp.Body.Close()

	imgBytes, _ := ioutil.ReadAll(resp.Body)
	fileName := `D:\迅雷下载\` + strconv.Itoa(int(time.Now().UnixNano())) + ".png"
	err = ioutil.WriteFile(fileName, imgBytes, 0644)
	if err == nil {
		fmt.Println(fileName, "download success")
	} else {
		fmt.Println(fileName, "download fail")
	}

}

func main() {

	var url = "http://www.2345.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("net error =", err)
		return
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)

	re := regexp.MustCompile(`<img.*?src="(\/\/www.*?)"`)
	rets := re.FindAllStringSubmatch(html, -1)
	for _, img := range rets {
		DownloadAsync("http:" + img[1])
	}
	downloadWG.Wait()
}
