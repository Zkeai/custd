package main

import (
	"fmt"
	"github.com/Zkeai/custd/app/config"
)

func Init_() error {
	err := config.Init()
	if err != nil {
		return fmt.Errorf("初始化失败：" + err.Error())
	}

	return nil
}
