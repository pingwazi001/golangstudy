package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

var (
	client *http.Client
)

func init() {
	client = &http.Client{}
}

func respResult(resp *http.Response, err error) (string, error) {
	//判断请求是否成功
	if err != nil {
		return "", err
	}
	//使用缓冲区去读取响应内容
	//定义一个缓冲区
	var readBuf [1024]byte
	result := bytes.NewBuffer(nil)
	for {
		len, err := resp.Body.Read(readBuf[0:])
		if err != nil && err == io.EOF {
			break //读取完成
		} else if err != nil {
			return "", err
		}
		result.Write(readBuf[0:len])
	}
	return result.String(), nil
}

func BenchmarkMain(b *testing.B) {
	c := make(chan string, b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			resp, err := client.Post("http://localhost:8080/post", "application/x-www-form-urlencoded", nil)
			str, err := respResult(resp, err)
			if err != nil {
				c <- err.Error()
				return
			}
			c <- str
		}()
	}
	for i := 0; i < b.N; i++ {
		b.Log(<-c)
	}
}
