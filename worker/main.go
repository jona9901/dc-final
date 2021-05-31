package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
    "strings"
    "encoding/json"

	pb "github.com/jona9901/dc-final/proto"
	"github.com/jona9901/dc-final/scheduler"
	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/req"
	"google.golang.org/grpc"

	// register transports
	_ "go.nanomsg.org/mangos/transport/all"
)

var (
	defaultRPCPort = 50051
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedFilterServer
}

var (
	controllerAddress = ""
	workerName        = ""
	tags              = ""
)

// Create channel to tell the controller a new woekload has been created
//var newWkld = make(chan scheduler.Workload)

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// SayHello implements helloworld.GreeterServer
func (s *server) ApplyFilter(ctx context.Context, in *pb.FilterRequest) (*pb.FilterReply, error) {
	log.Printf("RPC: Received: %v", in.GetFilter())
	return &pb.FilterReply{WorkloadID: "WorkloadID:  " + in.GetFilter()}, nil
}

// CreateWorkload implements filters.CreateWorkload
func (s *server) CreateWorkload(ctx context.Context, in *pb.WorkloadRequest) (*pb.WorkloadReply, error) {
	log.Printf("RPC: Creating workload with ID: %v", in.GetWorkloadID())
    log.Printf("Creating workload")
    workloadBuff := scheduler.Workload {
        WorkloadID: in.WorkloadID,
        Filter: in.Filter,
        WorkloadName: in.WorkloadName,
        Status: in.Status,
        RunningJobs: int(in.RunningJobs),
    }
//    newWkld <-workloadBuff
    client(workloadBuff)
	return &pb.WorkloadReply{Message: "Workload:  " + in.GetWorkloadName() + " scheduled."}, nil
}

func init() {
	flag.StringVar(&controllerAddress, "controller", "tcp://localhost:40899", "Controller address")
	flag.StringVar(&workerName, "worker-name", "hard-worker", "Worker Name")
	flag.StringVar(&tags, "tags", "gpu,superCPU,largeMemory", "Comma-separated worker tags")
}

// joinCluster is meant to join the controller message-passing server
/*
func joinCluster(newWorkload chan scheduler.Workload) {
	var sock mangos.Socket
	var err error
//	var msg []byte

	if sock, err = pub.NewSocket(); err != nil {
		die("can't get new pub socket: %s", err.Error())
	}

	log.Printf("Connecting to controller on: %s", controllerAddress)
	if err = sock.Dial(controllerAddress); err != nil {
		die("can't dial on pub socket: %s", err.Error())
	}
/ *
	// Empty byte array effectively subscribes to everything
	err = sock.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil {
		die("cannot subscribe: %s", err.Error())
	}* /
	for {
		if err = sock.Send([]byte("hola desde worker")); err != nil {
			die("Cannot recv: %s", err.Error())
		}
        / *if nWorkload scheduler.Workload := <-newWorkload
		if err = sock.Send([]byte(nWorkload)); err != nil {
			die("Failed publishing: %s", err.Error())
		}* /
	}
}
*/

func registerWorker(rpcPort int) {
    var sock mangos.Socket
    var err error
    var msg []byte

    worker := scheduler.Worker {
        WorkerName: workerName,
        Tags: strings.Split(tags, ","),
        WorkerPort: rpcPort,
    }

    workerBuff, err := json.Marshal(worker)
    if err != nil {
        die("Json error: %s", err.Error())
    }

	if sock, err = req.NewSocket(); err != nil {
		die("can't get new req socket: %s", err.Error())
	}
	if err = sock.Dial(controllerAddress); err != nil {
		die("can't dial on req socket: %s", err.Error())
	}
	log.Printf("Worker sending registering info")
	if err = sock.Send(workerBuff); err != nil {
		die("can't send message on push socket: %s", err.Error())
	}
	if msg, err = sock.Recv(); err != nil {
		die("can't receive date: %s", err.Error())
	}
    log.Printf("Message recieved from controller: %s", string(msg))
}

func client(newWorkload scheduler.Workload) {
	var sock mangos.Socket
	var err error
	var msg []byte

//    wkldbuff := <-newWorkload
    workloadJson, err := json.Marshal(newWorkload)

    if err != nil {
        die("Json error: %s", err.Error())
    }

	if sock, err = req.NewSocket(); err != nil {
		die("can't get new req socket: %s", err.Error())
	}
	if err = sock.Dial(controllerAddress); err != nil {
		die("can't dial on req socket: %s", err.Error())
	}
	log.Printf("Worker sending workload info")
	if err = sock.Send(workloadJson); err != nil {
		die("can't send message on push socket: %s", err.Error())
	}
	if msg, err = sock.Recv(); err != nil {
		die("can't receive date: %s", err.Error())
	}
    log.Printf("Message recieved from controller: %s", string(msg))
}

func getAvailablePort() int {
	port := defaultRPCPort
	for {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			port = port + 1
			continue
		}
		ln.Close()
		break
	}
	return port
}

func main() {
	flag.Parse()

	// Subscribe to Controller
//    newWorkload := make(chan scheduler.Workload)
	//go joinCluster(newWorkload)
//    go client(newWorkload)
    //go client()

	// Setup Worker RPC Server
	rpcPort := getAvailablePort()
    registerWorker(rpcPort)
	log.Printf("Starting RPC Service on localhost:%v", rpcPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", rpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFilterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
//    wkBuff := <-newWkld
//    newWorkload <-wkBuff
    // Close channels
    //close(newWkld)
}
