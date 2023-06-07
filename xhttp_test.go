/*
 * Author: DG9Jww
 * Date: 2023/3/10
 */

package xhttp

import (
	"fmt"
	"log"
	"testing"
)

func TestDoA(t *testing.T) {
	cli := NewClient()
	cli.SetHeaders(map[string]string{
		"aaa": "bbbb",
		"bbb": "bbbb",
		"ccc": "bbbb",
	})
	cli.SetProxy("http://127.0.0.1:8080")
	resp, err := cli.DoA("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}

func TestDoB(t *testing.T) {
	cli := NewClient()
	cli.SetHeaders(map[string]string{
		"aaa": "bbbb",
		"bbb": "bbbb",
		"ccc": "bbbb",
	})
	cli.SetProxy("http://127.0.0.1:8080")
	resp, err := cli.DoB("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}

func TestDoC(t *testing.T) {
	cli := NewClient()
	cli.SetHeaders(map[string]string{
		"aaa": "bbbb",
		"bbb": "bbbb",
		"ccc": "bbbb",
	})
	cli.SetProxy("http://127.0.0.1:8080")
	resp, err := cli.DoC("GET", "http://www.baidu.com", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Header)
}
