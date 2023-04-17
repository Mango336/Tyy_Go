package main

import (
	"TyyGin/model"
	"TyyGin/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	EG errgroup.Group
	// WG sync.WaitGroup
)

func main() {
	fmt.Println("Gin Web is loading...")
	model.InitDb()
	routes.InitRouter()
	// Example()
	// SyncRouter()
}
func SyncRouter() {
	server01 := &http.Server{
		Addr:         ":8081",
		Handler:      routes.Router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server02 := &http.Server{
		Addr:         ":8082",
		Handler:      routes.Router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// 方法一：借助errgroup.Group分别启动两个服务
	EG.Go(func() error {
		return server01.ListenAndServe()
	})
	EG.Go(func() error {
		return server02.ListenAndServe()
	})
	if err := EG.Wait(); err != nil {
		log.Fatal(err)
	}

	// 方法二：使用两个goroutine开启服务
	// WG.Add(2)
	// go func() {
	// 	defer WG.Done()
	// 	server01.ListenAndServe()
	// }()
	// go func() {
	// 	defer WG.Done()
	// 	server02.ListenAndServe()
	// }()
	// WG.Wait()
}
