package controller

import (
	"fmt"
	"log"
	"os"
	"time"
    "encoding/json"
    "strings"

    "github.com/jona9901/dc-final/scheduler"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/rep"
	//"go.nanomsg.org/mangos/protocol/req"

	// register transports
	_ "go.nanomsg.org/mangos/transport/all"
)

var controllerAddress = "tcp://localhost:40899"

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

var availableWorkers = []scheduler.Worker{}
var availableWorkloads = []scheduler.Workload{}

/*
func server() {
	var sock mangos.Socket
	var err error
    var msg []byte

	if sock, err = pub.NewSocket(); err != nil {
		die("can't get new pub socket: %s", err)
	}

	if err = sock.Listen(controllerAddress); err != nil {
		die("can't listen on pub socket: %s", err.Error())
	}

	for {
		// Could also use sock.RecvMsg to get header
		d := date()
		log.Printf("Controller: Publishing Date %s\n", d)

        if msg, err = sock.Recv(); err != nil {
            die("Failed recv: %s", err.Error())
        }

        log.Printf("Recv: %s", string(msg))

		if err = sock.Send([]byte(d)); err != nil {
			die("Failed publishing: %s", err.Error())
        }
	}
}
*/

func socketServer(avlbWorkers chan []scheduler.Worker, avlbWorkloadList chan []scheduler.Workload) {
    var sock mangos.Socket
    var err error
    var msg []byte

    log.Printf("Initiating Socket")

	if sock, err = rep.NewSocket(); err != nil {
		die("can't get new rep socket: %s", err)
	}

    if err = sock.Listen(controllerAddress); err != nil {
		die("can't listen on rep socket: %s", err.Error())
	}

    for {
        msg, err = sock.Recv()
        if err != nil {
            die("Cannot receive on rep socket: %s", err.Error())
        }
        log.Printf("Message: %s", string(msg))

        if strings.Contains(string(msg), "WorkerName") {
            newWorker := scheduler.Worker{}
            json.Unmarshal(msg, &newWorker)
            log.Printf(newWorker.WorkerName)
            availableWorkers = append(availableWorkers, newWorker)
            avlbWorkers <-availableWorkers
        }
        if strings.Contains(string(msg), "WorkloadID") {
            newWorkload := scheduler.Workload{}
            json.Unmarshal(msg, &newWorkload)
            availableWorkloads = append(availableWorkloads, newWorkload)
            avlbWorkloadList <-availableWorkloads
            log.Printf("New workload created: %s", newWorkload.WorkloadID)
        }

		d := date()
        log.Printf("Controller: responding to worker at date: %s\n", d)
		if err = sock.Send([]byte(d)); err != nil {
			die("Failed sublishing: %s", err.Error())
		}
    }
}

func Start(avlbWorkerList chan []scheduler.Worker, avlbWorkloadList chan []scheduler.Workload) {
    log.Printf("Controller running")
    //go run server()
    go socketServer(avlbWorkerList, avlbWorkloadList)
//    getWorker(avlbWorkerList)
    //go run client()
    /*
	var sock mangos.Socket
	var err error
    var msg []byte
	if sock, err = sub.NewSocket(); err != nil {
		die("can't get new sub socket: %s", err)
	}

	if err = sock.Listen(controllerAddress); err != nil {
		die("can't listen on sub socket: %s", err.Error())
	}

    if err = sock.Dial(controllerAddress); err != nil {
		die("can't dial on sub socket: %s", err.Error())
    }
    // Empty byte array effectively subscribes to everything
	//err = sock.SetOption(mangos.OptionSubscribe, []byte(""))
	//if err != nil {
//		die("cannot subscribe: %s", err.Error())
//	}

	for {
		// Could also use sock.RecvMsg to get header
		d := date()
		log.Printf("Controller: Publishing Date %s\n", d)
		if err = sock.Send([]byte(d)); err != nil {
			die("Failed sublishing: %s", err.Error())
		}
        for {
            if msg, err = sock.Recv(); err != nil {
                die("Cannot Recv: %s", err.Error())
            }
            log.Printf("Message from worker- %s", string(msg))
        }
		time.Sleep(time.Second * 3)
	}
    */
}
