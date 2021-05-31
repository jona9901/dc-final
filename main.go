package main

import (
	//"fmt"
	"log"
	//"math/rand"
//	"time"

	"github.com/jona9901/dc-final/api"
	"github.com/jona9901/dc-final/controller"
	"github.com/jona9901/dc-final/scheduler"
)

func main() {
	log.Println("Welcome to the Distributed and Parallel Image Processing System")

    // Create all channels
	jobs := make(chan scheduler.Job)
	workloads := make(chan scheduler.Workload)
//    availableWorkers := make(chan []scheduler.Worker)
//    availableWorkloads := make(chan []scheduler.Workload)

	// API
//    newJob := make(chan scheduler.Job)
	go api.Start(workloads)

	// Start Controller
	go controller.Start()//availableWorkers, availableWorkloads)

	// Start Scheduler
	go scheduler.Start(workloads)//, availableWorkers)
	// Send sample jobs
	//sampleJob := scheduler.Job{Address: "localhost:50051", RPCName: "hello"}

 //   nJob := scheduler.Job{}
    for {
  /*      nJob = <-jobs
        newJob <-nJob*/
    }
    /*
	for {
		sampleJob.RPCName = fmt.Sprintf("hello-%v", rand.Intn(10000))
		jobs <- sampleJob
		time.Sleep(time.Second * 5)
	}
    */
    close(jobs)
    close(workloads)
    /*
    close(availableWorkers)
    close(availableWorkloads)
    */
}
