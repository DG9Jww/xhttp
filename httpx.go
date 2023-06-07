/*
 * Author: DG9Jww
 * Date: 2023/3/10
 */

package xhttp

import (
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
	headers    http.Header
	proxy      string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		headers: make(http.Header),
	}
}

// 设置头部
func (c *Client) SetHeader(key, value string) {
	c.headers.Set(key, value)
}

// 设置多个头部
func (c *Client) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		c.SetHeader(key, value)
	}
}

// 设置代理
func (c *Client) SetProxy(proxyURL string) {
	c.proxy = proxyURL
}

/*******
判断是否压缩
*******/

// 尝试解压缩并读取body
func DecompressAndReadBody(resp *http.Response) ([]byte, error) {
	var reader io.ReadCloser
	var err error
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
	case "deflate":
		reader = flate.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

/*******
发起请求
*******/

// 返回请求体string
func (c *Client) DoA(method string, host string, body []byte) (string, error) {
	req, err := http.NewRequest(method, host, strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	req.Header = c.headers

	if c.proxy != "" { // check if proxy URL is set
		proxyURL, err := url.Parse(c.proxy)
		if err != nil {
			return "", err
		}
		c.httpClient.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := DecompressAndReadBody(resp)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// 返回请求体[]byte
func (c *Client) DoB(method string, host string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, host, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	req.Header = c.headers

	if c.proxy != "" { // check if proxy URL is set
		proxyURL, err := url.Parse(c.proxy)
		if err != nil {
			return nil, err
		}
		c.httpClient.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := DecompressAndReadBody(resp)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

// 返回*http.Response
func (c *Client) DoC(method string, host string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, host, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	req.Header = c.headers

	if c.proxy != "" { // check if proxy URL is set
		proxyURL, err := url.Parse(c.proxy)
		if err != nil {
			return nil, err
		}
		c.httpClient.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
