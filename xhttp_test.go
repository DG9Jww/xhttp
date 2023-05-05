/*
 * Author: DG9Jww
 * Date: 2023/3/10
 */

package xhttp

import (
	"fmt"
	"testing"
)

func TestGetString(t *testing.T) {
	cli := NewClient()
	s, err := cli.GetString("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestPostString(t *testing.T) {
	cli := NewClient()
	cli.SetProxy("http://127.0.0.1:8080")
	s, err := cli.PostString("http://www.baidu.com", "888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)

	cli.PostString("http://www.baidu.com", "666")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
