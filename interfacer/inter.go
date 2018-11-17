//在 Golang 中，interface 是一组 method 的集合，是 duck-type programming 的一种体现。不关心属性（数据），只关心行为（方法）。

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type MyInterface interface {
	Print(s string) string
}

type HttpGeter interface {
	// Get method
	Get(url string) string
}

type HttpPoster interface {
	// Post method
	Post(url string, params map[string]string) string
}

// 接口的组合
type Httper interface {
	HttpGeter
	HttpPoster
}

// 调用接口方式
func TestFunc(x MyInterface) {

}

// 实际调用接口方法
func Get(hg HttpGeter) string {
	content := hg.Get("http://www.baidu.com")
	//fmt.Print(content)
	return content
}

// MyStruct实现了MyInterface接口
// MyStruct同时也实现了Httper方法
type MyStruct struct {
	Content string
}

func (me *MyStruct) Print(s string) string {
	fmt.Println("Show s: %s", s)
	return s
}

func (me *MyStruct) Get(url string) string {

	res, err := http.Get(url)
	if err != nil {
		log.Printf("Http request fail, %s\n", err)
	}

	content := me.printContent(res.Body)
	me.Content = content
	return content
}

func (me *MyStruct) printContent(r io.ReadCloser) string {
	defer r.Close()
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("Read fail, %s\n", err)
		return ""
	}
	return string(bs)
}

func (me *MyStruct) Post(url string, params map[string]string) string {
	return ""
}

func main() {
	var me MyStruct
	me.Print("Nice Job.")
	// 因为MyStruct实现了MyInterface接口，所以TestFunc可以接收me作为参数
	TestFunc(&me)

	// me 同时实现了HttpGeter接口，而Get接收Get方法，所以可以直接调用
	Get(&me)
	// 不能直接赋值，必须通过指针来传值
	fmt.Println(me.Content)
}
