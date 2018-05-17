package main

import (
	"fmt"
	"regexp"
)

func main() {
	//	match, _ := regexp.MatchString(`<img[^>]*?(?<!_)src="(?!http://cdn.17zuoye.com)`, `<img src="http://cdn.17zuoye.com>`)
	match, _ := regexp.MatchString(`<img[^>]*?src="[^"]+"`, `<img _src="http://cdn.17zuoye.com">`)
	fmt.Println(match)

	match, _ = regexp.MatchString(`<img[^>]*? src="[^"]+"`, `<img _src="http://cdn.17zuoye.com">`)
	fmt.Println(match)

	match, _ = regexp.MatchString(`<img[^>]*src="(http://cdn\.17zuoye\.com)`, `<img src="http://www.baidu.com">`)
	fmt.Println(match)

	//	re := regexp.MustCompile(`</mark>(_+)`)
	//	fmt.Println(re.FindStringSubmatch(`</mark>____+__`))
}

// <img src
