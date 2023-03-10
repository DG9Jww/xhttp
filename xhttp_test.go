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
	s, err := GetString("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestPostString(t *testing.T) {
	s, err := PostString("http://www.baidu.com", 50)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
