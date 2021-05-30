package scheduler

import (
	"context"
	"log"
	"time"

	pb "github.com/jona9901/dc-final/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type Workload struct {
	WorkloadID	string
	Filter		string
	WorkloadName	string
	Status		string
	ImageId		string
	RunningJobs	int
	filtered_images	[50] string
}

type Job struct {
	Address string
	RPCName string
}

func schedule(job Job, workload Workload) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFilterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
//	r, err := c.ApplyFilter(ctx, &pb.FilterRequest{WorkloadID: workloadID,})
    r, err := c.CreateWorkload(ctx, &pb.WorkloadRequest{
        WorkloadID: workload.WorkloadID,
        Filter: 0,
        WorkloadName: workload.WorkloadName,
        Status: 0,
        running_jobs: 0,
    })
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
//	log.Printf("Scheduler: RPC respose from %s : %s", address, r.GetWorkloadID())
	log.Printf("Scheduler: RPC respose from %s : %s", address, r.GetMessage())
}

func Start(jobs chan Job, workloads chan Workload) error {
	for {
		job := <-jobs
		workload := <-workloads
		schedule(job, workload)
	}
	return nil
}
