package main

import (
	"event/app/app"
	"event/config"
	"flag"
	"fmt"
)

var configFlag = flag.String("config", "./config.toml", "toml env file not found") // 실행시 flag값을 추가해서 값을 받아올 수 있게 해줌

func main() {
	flag.Parse()
	a := app.NewApp(config.NewConfig(*configFlag))
	fmt.Println(a)
}
