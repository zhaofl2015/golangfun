package main

import (
	"fmt"
	"question-bank/utils"
	"regexp"
	//	"strconv"
	//	"regexp"
)

//func FilterFullWidthCharacters(text string) string {
//	ascii_str := strconv.QuoteToASCII(text)
//	re := regexp.MustCompile(`(\\u.{4})`)
//	ascii_filter := re.ReplaceAllString(ascii_str, "")
//	return ascii_filter
//}

func main() {
	//	re := regexp.MustCompile(`(\\u.{4})`)
	//	s := re.ReplaceAllString("\u554a", " ")
	//	fmt.Println(s)
	//	quoted := strconv.QuoteRuneToASCII('啊') // quoted = "'\u554a'"
	//	unquoted := quoted[1 : len(quoted)-1]   // unquoted = "\u554a"
	//	fmt.Println(unquoted)
	//	s = re.ReplaceAllString(unquoted, " ")
	//	fmt.Println(s)
	//	fmt.Println("\u554a" == "啊")

	//	for
	//	fmt.Println(strconv.QuoteToASCII("哈哈。你好，今天天气·不太好a,真的吗’？are you serious"))
	fmt.Println(utils.FilterFullWidthCharacters("哈哈。—你好，今天天气·不太好a,真的吗’\\n？are you serious"))
	//	fmt.Println(strconv.QuoteToGraphic("哈哈。你好，今天天气·不太好a,真的吗？are you serious"))
	//	fmt.Println(strconv.CanBackquote(strconv.QuoteToASCII("哈哈。你好，今天天气·不太好a,真的吗？are you serious")))
	fmt.Println(len(""))

	//	fmt.Println(regexp.MatchString(`[\u4e00-\u9fa5]{3,8}`, "今天去哪里"))
	//	fmt.Println(regexp.MatchString(`[\u4e00-\u9fa5]{3,8}`, strconv.QuoteToASCII("今天去哪里")))

	re := regexp.MustCompile("[\u4e00-\u9fa5]+")
	fmt.Println(re.MatchString("111我1牛逼啊1"))
}
