package utils

import (
	"fmt"
	"io"
	"os"
	"time"
)

func CheckFile(Filename string) bool {
	var exist = true
	if _, err := os.Stat(Filename); os.IsNotExist(err) {
		exist = false
		if err != nil {
			fmt.Println("not found log")
		}
	}
	return exist
}

//写入文件
func Logfile(logType string, log string) {
	var f1 *os.File
	var err1 error

	filenames := "./file/logs/" + time.Now().Format("20060102") + ".log" //也可将name作为参数传进来

	if CheckFile(filenames) { //如果文件存在
		f1, err1 = os.OpenFile(filenames, os.O_APPEND|os.O_WRONLY, 0666) //打开文件,第二个参数是写入方式和权限
		if err1 != nil {
			fmt.Println("文件存在，已打开")
		}
	} else {
		f1, err1 = os.Create(filenames) //创建文件
		if err1 != nil {
			fmt.Println("创建文件失败")
		}
	}
	_, err1 = io.WriteString(f1, logType+time.Now().Format("2006-01-02 15:04:05")+log+"\n") //写入文件
	if err1 != nil {
		fmt.Println(err1)
	}
	return
}
