// log
package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// log句柄对象
type Handle struct {
	FirstTag string      //输出标记字符
	Logger   *log.Logger //logger对象
	File     *os.File    //log对应的文件对象
	sign     string      //左边字符标记
}

func (s *Handle) Init(folderName string, firstTag string, sign string) {

	var err error
	//创建logdata文件夹
	_, err = os.Stat("logdata")
	if err != nil {
		os.MkdirAll("logdata", 0777)
	}

	//创建本服的log文件夹
	logFolder := "logdata/" + folderName
	_, err = os.Stat(logFolder)
	if err != nil {
		os.MkdirAll(logFolder, 0777)
	}

	//获取当前时间
	dataTime := time.Now()

	//构建文件路径
	filePath := fmt.Sprintf(logFolder+"/%d_%02d_%02d_%s.txt",
		dataTime.Year(), dataTime.Month(), dataTime.Day(),
		firstTag)

	var file *os.File
	var bExitst bool

	//判断文件夹是否存在，不存在就创建一个

	//判断文件是否存在
	bExitst = Exist(filePath)
	if bExitst == true {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0)
		if err != nil {
			fmt.Println(err)
			fmt.Println("打开", filePath, "失败")
			return
		}

	} else {
		file, err = os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			fmt.Println("创建", filePath, "失败")
			return
		}
	}

	s.File = file
	s.FirstTag = firstTag
	s.sign = sign
	s.Logger = log.New(file, "", log.LstdFlags|log.Lmicroseconds)
}

func (s *Handle) Log(v ...interface{}) {

	//获取函数调用文件及调用行数
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}

	//获取短文件名
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}

	s.Logger.Println(s.sign, s.FirstTag, s.sign, fmt.Sprint(v...), "[file:", short, "line:", line, "]")
}
