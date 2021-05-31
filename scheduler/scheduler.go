package scheduler

import (
	"context"
	"log"
	"time"
	"math/rand"

	pb "github.com/jona9901/dc-final/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type Workload struct {
	WorkloadID	    string
	Filter		    string
	WorkloadName	string
	Status		    string
	RunningJobs	    int
    ImageID         string
    WorkerAddress   string
    FilteredImages  []string
}

type Job struct {
	Address string
	RPCName string
}

type Worker struct {
    WorkerName      string
    WorkerPort      int
    Tags            []string
}


func schedule(workload Workload) { //, worker Worker) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFilterClient(conn)

    if workload.WorkerAddress == "" {
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
    //	r, err := c.ApplyFilter(ctx, &pb.FilterRequest{WorkloadID: workloadID,})
        r, err := c.CreateWorkload(ctx, &pb.WorkloadRequest{
            WorkloadID: workload.WorkloadID,
            Filter: workload.Filter,
            WorkloadName: workload.WorkloadName,
            Status: workload.Status,
            RunningJobs: uint64(workload.RunningJobs),
        })
        if err != nil {
            log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Scheduler: RPC respose from %s : %s", address, r.GetMessage())
    } else {
        log.Printf("hola")

        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.ApplyFilter(ctx, &pb.FilterRequest {
            Filter: workload.Filter,
            ImageID: workload.ImageID,
            WorkloadID: workload.WorkloadID,
        })
        if err != nil {
            log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Scheduler: RPC respose from %s : %s", address, r.GetMessage())
    }
}


func Start(workloads chan Workload) {//jobs chan Job, workloads chan Workload) { //, availableWorkers chan []Worker) error {
	for {
//        if (Job{}) != jobs
		//job := <-jobs
		workload := <-workloads
        //workers := <-availableWorkers
        rand.Seed(time.Now().Unix())
        //worker := workers[rand.Intn(len(workers))]
		schedule(workload) //, worker)
        workloads <-workload
//        schedule2(job, workload)
	}
	//return nil
}
