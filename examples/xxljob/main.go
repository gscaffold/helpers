package main

import (
	"context"
	"log"
	"time"

	"github.com/gscaffold/helpers/xxljob"
	"github.com/xxl-job/xxl-job-executor-go"
)

func main() {
	client, err := xxljob.New(
		xxljob.ServerAddr("http://127.0.0.1:8666/xxl-job-admin"),
		xxljob.AccessToken("default_token"),
		xxljob.ExecutorIP("10.0.0.1"),
		xxljob.ExecutorPort(9999),
		xxljob.RegistryKey("example"),
	)
	if err != nil {
		log.Fatal(err)
	}
	client.RegTask("gscaffold", func(cxt context.Context, param *xxl.RunReq) string {
		log.Printf("gscaffold example start, param:%+v\n", param)
		time.Sleep(time.Second)
		log.Printf("gscaffold example done.")
		return "task done"
	})
	err = client.Run()
	if err != nil {
		log.Fatal(err)
	}
}
