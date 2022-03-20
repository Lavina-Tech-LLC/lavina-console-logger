package llog

import "fmt"

const (
	infoF    = "\033[1;34m[INFO]\033[0m %s"
	noticeF  = "\033[1;36m[Notice]\033[0m %s"
	warningF = "\033[1;33m[Warning]\033[0m %s"
	errorF   = "\033[1;31m[Error]\033[0m %s"
	debugF   = "\033[0;36m[Debug]\033[0m %s"
)

func Info(s interface{}) {
	fmt.Printf(infoF, s)
	fmt.Println("")
}

func Notice(s interface{}) {
	fmt.Printf(noticeF, s)
	fmt.Println("")
}

func Warning(s interface{}) {
	fmt.Printf(warningF, s)
	fmt.Println("")
}

func Error(s interface{}) {
	fmt.Printf(errorF, s)
	fmt.Println("")
}

func Debug(s interface{}) {
	fmt.Printf(debugF, s)
	fmt.Println("")
}
