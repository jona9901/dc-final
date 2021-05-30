package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jona9901/dc-final/api"
	"github.com/jona9901/dc-final/controller"
	"github.com/jona9901/dc-final/scheduler"
)

func main() {
	log.Println("Welcome to the Distributed and Parallel Image Processing System")

	// API
	jobs := make(chan scheduler.Job)
	workloads := make(chan scheduler.Workload)
	go api.Start(workloads)

	// Start Controller
	go controller.Start()

	// Start Scheduler
	go scheduler.Start(jobs, workloads)
	// Send sample jobs
	sampleJob := scheduler.Job{Address: "localhost:50051", RPCName: "hello"}

	for {
		sampleJob.RPCName = fmt.Sprintf("hello-%v", rand.Intn(10000))
		jobs <- sampleJob
		time.Sleep(time.Second * 5)
	}
}
