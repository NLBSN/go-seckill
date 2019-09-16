package main

import (
	"github.com/astaxie/beego"
	_ "SecKill/SecAdmin/router"
	"fmt"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}
