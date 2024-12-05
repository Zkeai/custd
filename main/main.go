package main

import (
	"fmt"
)

func main() {
	if err := Init_(); err != nil {

		panic(err)
	}

	fmt.Println("sixusdt 启动成功，当前版本：" + "1.0.0")

}
