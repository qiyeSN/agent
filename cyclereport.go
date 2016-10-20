package main

import (
	"github.com/elastic/gosigar/get"
	"github.com/robfig/cron"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "protobuf/report"
	//hs "protobuf/report/hostMessage"
	//"os"
)

const (
	address = "localhost:50051"
	//defaultVersion = "1"
)

func main() {
	spec := "0/30, *, *, *, ?, *"
	c := cron.New()
	c.AddFunc(spec, report)
	//if err != nil {
	//	fmt.Println("some wrong")
	//	return
	//}
	c.Start()
	select {}
}

func callYourFunc() {
	//a := 0
	log.Println("callYourFunc come here.")
	//fmt.Println("heheheh", 1/a)
}

func report() {
	hostMess := get.GetHost()
	processMess := get.GetProcess()
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	//version := defaultVersion
	//if len(os.Args) > 1 {
	//	version = os.Args[1]
	//}
	r, err := c.SendReport(context.Background(), &pb.Report{Host: &hostMess, Process: processMess})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %d", r.Mess)
}
