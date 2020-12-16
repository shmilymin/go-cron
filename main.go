package main

import (
	"go-cron/routers"
)

func main() {
	r := routers.InitRouters()
	r.Run(":8888")
}
