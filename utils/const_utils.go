package utils

func GetVisibleText(visible_value int) string {
	var VisibleDict map[int]string
	VisibleDict = make(map[int]string)
	VisibleDict[1] = "所有人可见"
	VisibleDict[2] = "登录可见"
	VisibleDict[3] = "作者可见"
	res, _ := VisibleDict[visible_value]
	return res
}
