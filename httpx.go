/*
 * Author: DG9Jww
 * Date: 2023/3/10
 */

package xhttp

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	cli *http.Client
	req *http.Request
}

func NewClient() *Client {
	r, _ := http.NewRequest("GET", "", nil)
	c := Client{
		cli: &http.Client{},
		req: r,
	}
	return &c
}

// 设置代理, 会覆盖掉 http.Transport
func (c *Client) SetProxy(proxy string) {
	p := func(*http.Request) (*url.URL, error) {
		return url.Parse(proxy)
	}
	c.cli = &http.Client{Transport: &http.Transport{Proxy: p}}
}

// 设置超时
func (c *Client) SetTimeout(t time.Duration) {
	c.cli.Timeout = t
}

// 设置忽略证书 ,会覆盖掉 http.Transport
func (c *Client) SkipVerify() {
	c.cli = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
}

// 设置请求头，会覆盖掉原来的值
func (c *Client) SetHeader(k string, v string) {
	c.req.Header.Set(k, v)
}

// 新添请求头部
func (c *Client) AddHeader(k string, v string) {
	c.req.Header.Add(k, v)
}

//-----------------------------------------------------------------------

// GET获取响应体字符串
func (c *Client) GetString(url string) (string, error) {
	err := setURL(url, c.req)
	if err != nil {
		return "", err
	}

	resp, err := c.cli.Do(c.req)
	if err != nil {
		return "", err
	}
	body := resp.Body
	defer body.Close()
	b, _ := io.ReadAll(body)
	return string(b), nil
}

// POST获取响应体字符串
// body为string或[]byte
func (c *Client) PostString(url string, body interface{}) (string, error) {
	err := setURL(url, c.req)
	err = setMethod("POST", c.req)
	if err != nil {
		return "", err
	}

	c.setPostBody(body)
	resp, err := c.cli.Do(c.req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return string(b), nil
}

// GET获取 *http.Response
func (c *Client) GetResponse(url string) (*http.Response, error) {
	err := setURL(url, c.req)
	if err != nil {
		return nil, err
	}
	resp, err := c.cli.Do(c.req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// POST获取 *http.Response
func (c *Client) PostResponse(url string, body interface{}) (*http.Response, error) {
	err := setURL(url, c.req)
	err = setMethod("POST", c.req)
	if err != nil {
		return nil, err
	}

	c.setPostBody(body)
	resp, err := c.cli.Do(c.req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 设置请求体，只能是string/[]byte/nil
func (c *Client) setPostBody(body interface{}) {

	/*******
	如果body是nil
	*******/
	if body == nil {
		c.req.Body = nil

		/*******
		如果是string
		*******/
	} else if tmp, ok := body.(string); ok {
		c.req.Body = io.NopCloser(strings.NewReader(tmp))

		/*******
		如果是[]byte
		*******/
	} else if tmp2, ok := body.([]byte); ok {
		c.req.Body = io.NopCloser(bytes.NewReader(tmp2))
	}

}

// 设置url
func setURL(u string, req *http.Request) error {
	h, err := url.Parse(u)
	if err != nil {
		return err
	}
	req.URL = h
	return nil
}

var validMethod = map[string]bool{
	"GET":     true,
	"HEAD":    true,
	"POST":    true,
	"PUT":     true,
	"DELETE":  true,
	"TRACE":   true,
	"CONNECT": true,
}

/*
设置请求方法
支持的请求方法如下

	"OPTIONS"
	"GET"
	"HEAD"
	"POST"
	"PUT"
	"DELETE"
	"TRACE"
	"CONNECT"
*/
func setMethod(method string, req *http.Request) error {
	//非法请求方法
	if _, ok := validMethod[method]; !ok {
		return errors.New("Invalid request method!")
	}
	req.Method = method
	return nil
}
