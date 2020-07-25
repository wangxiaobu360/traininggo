package main

import (
	"os/exec"
	"fmt"
	"time"
)

func GetCode(workSpace string)  {
	command := fmt.Sprintf("cd %s && git pull", workSpace)
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(string(out))
	}
}

func GetAdminMain(srcMain string, dstMain string) {
	command := fmt.Sprintf("go build -o %s %s", dstMain, srcMain)
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(string(out))
	}
}

func GetNowTag() string{
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	//输出格式可以自定义
	return fmt.Sprintf("%d_%02d_%02d_%02d_%02d_%02d", year, month, day, hour, minute, second)
}

func DoShell(workSpace string) {
	nowTime := GetNowTag()
	command := fmt.Sprintf("cd %s && ./do.sh %s", workSpace, nowTime)
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(string(out))
	}
}

func main() {
	fmt.Println("hello")
}
