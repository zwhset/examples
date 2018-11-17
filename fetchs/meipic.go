package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func WriteFile(title, url string) {

	path := "fetchs/images"
	os.MkdirAll(path, 0644)
	filename := filepath.Join(path, title) + ".jpg"

	if PathExists(filename) {
		i := 1
		for {
			title = title + strconv.Itoa(i)

			filename = filepath.Join(path, title) + ".jpg"
			log.Printf(filename)
			if !PathExists(filename) {
				break
			}
			i++
		}
	}

	//url := <-urls
	log.Printf("Start Download %s -> %s\n", url, filename)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Download Pic %s Faild, %s\n", url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read Download Pic %s Faild, %s\n", url, err)
	}

	fd, err := os.Create(filename)
	if err != nil {
		log.Printf("Create file %s Faild, %s\n", filename, err)
	}
	_, err = io.Copy(fd, bytes.NewReader(body))
	if err != nil {
		log.Printf("Write File %s Faild, %s\n", filename, err)
	}

	log.Printf("End Download %s -> %s\n", url, filename)
}

// 文件是否存在，存在返回true
func PathExists(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func main() {
	getUrl := "http://www.27270.com/ent/meinvtupian/"

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"

	c.AllowedDomains = []string{"www.27270.com", "t2.hddhhn.com", "t1.hddhhn.com", "t3.hddhhn.com"}

	// 1级
	c.OnHTML(".NewPages a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if !strings.HasPrefix(link, getUrl) {
			link = filepath.Join(getUrl, link)
		}
		//fmt.Printf("Top Page found: %q -> %s\n", e.Text, link)

		c.Visit(e.Request.AbsoluteURL(link))
	})

	// 2级页
	c.OnHTML(".MeinvTuPianBox a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if !strings.HasPrefix(link, getUrl) {
			link = filepath.Join(getUrl, link)
		}

		//fmt.Printf("Laver 2 Page found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// 最终页
	c.OnHTML(".articleV4Body  img[src]", func(e *colly.HTMLElement) {

		src := e.Attr("src")
		//title := e.Attr("alt")
		_, title := filepath.Split(src)
		title = strings.Split(title, ".")[0]

		fmt.Printf("Pic found: %s -> %s\n", title, src)
		//os.Exit(1)

		go WriteFile(title, src)

	})

	//c.OnResponse(func(resp *colly.Response) {
	//	//fmt.Println("onResponse: ", string(resp.Body))
	//})

	c.Visit(getUrl)

	c.Wait()

	time.Sleep(time.Minute * 10)

}
