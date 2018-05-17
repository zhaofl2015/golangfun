package main

import (
	"fmt"
)

// 全角转半角
func Q2B(text string) string {
	var restext string
	restext = ""

	for _, ch := range text {
		if ch == 0x3000 {
			restext += fmt.Sprintf("%c", 32)
		} else if ch >= 0xFF01 && ch <= 0xFF5E {
			ch -= 0xfee0
			restext += fmt.Sprintf("%c", ch)
		} else {
			restext += fmt.Sprintf("%c", ch)
		}
	}

	return restext
}

func main() {
	text := Q2B("哈哈，a。，，，,你好啊")
	fmt.Println(text)
}
