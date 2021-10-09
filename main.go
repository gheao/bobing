package main

import (
	"t1/routes"
)

// 返回数据的结构

func main() {
	r := routes.ServerInit()

	r.Run(":8080")
}
