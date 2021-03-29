package main

import (
	"log"
	"port-scanner/conn"
	"port-scanner/pool"
	"strconv"
	"sync"
)

func main() {
	var poolSize = 1000
	var wg sync.WaitGroup
	aPool := *pool.New(poolSize, &wg)

	for i := 1; i < 65535; i++ {
		wg.Add(1)
		aPool.Execute(checkHostOnPort, "localhost", i)
	}
	wg.Wait()
}

func checkHostOnPort(args ...interface{}) {
	host := args[0].(string)
	port := strconv.Itoa(args[1].(int))
	check := conn.Check("tcp", host, port)
	log.Printf("port: %s - %+v", port, check)
}
