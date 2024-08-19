package main

import (
	"github.com/tonnytg/kmetrics/pkg/runtime"
	"github.com/tonnytg/kmetrics/pkg/webserver"
	"log"
	"sync"
)

func main() {
	log.Println("Start Kmetrics")
	var wg sync.WaitGroup

	// Start Runtime to collect information about Kubernetes Cluster
	// This data will save in local sqlite3
	wg.Add(1)
	go func() {
		runtime.Start()
		wg.Done()
	}()

	// Start Webserver to see all information using HTML page
	// This pages can see using PortForward in Service
	wg.Add(1)
	go func() {
		webserver.Start()
		wg.Done()
	}()
	wg.Wait()
}
