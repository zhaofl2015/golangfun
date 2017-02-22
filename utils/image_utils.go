package utils

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"time"

	//	"github.com/astaxie/beego"
)

func GetRandomImageLocal(count int) []string {

	var ret_files []string

	files, err := ioutil.ReadDir(FirstImagePath)
	if err != nil {
		panic(errors.New("No Such File or Directory"))
	}

	for _, file := range files {
		ret_files = append(ret_files, FirstImagePath+"/"+file.Name())
	}

	rand.Seed(int64(time.Now().Second()))
	index_lst := rand.Perm(len(ret_files))[:count+1]

	var tmp []string

	for _, ind := range index_lst {
		tmp = append(tmp, ret_files[ind])
	}

	return tmp
}
