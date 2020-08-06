package main

import (
	"fmt"
	api "frame/api"
	"frame/util"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const ServiceName = "frame"
const ServiceTag = "grpc_tag"

func main() {
	rand.Seed(time.Now().UnixNano())
	grpc.NewServer()
	defaultPort := ":8081"
	badTaste := false
	port := defaultPort
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	if len(os.Args) == 3 {
		port = os.Args[1]
		badTaste = os.Args[2] == "1"
	}

	fmt.Printf("orgs(%v) %+v \n", len(os.Args), os.Args)
	grpcServer := grpc.NewServer()
	util.RegisterHealthCheck(grpcServer)
	api.RegisterSServiceServer(grpcServer, newSService(port))

	sw := &sync.WaitGroup{}
	sw.Add(1)
	if badTaste { //8081作为坏节点场景，随机崩坏
		//go util.ListenConsulStats(ServiceName, ServiceTag) //打印监控状态
		startBadTasteGrpc(grpcServer, port, sw)
	} else {
		startGrpc(grpcServer, port, sw)
	}
	sw.Wait()
	fmt.Printf("\n==========finish listen service(%v) ==========\n", ServiceName)

}

func startGrpc(grpcServer *grpc.Server, port string, sw *sync.WaitGroup) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen port(%v): %v", port, err)
		sw.Done()
		return
	}

	fmt.Printf("start listen port(%v)\n", port)
	go func() {
		errr := util.RegisterGrpcServiceToConsul(&util.RegisterGrpcServiceArgs{
			Name:    ServiceName,
			Port:    util.StrToInt(strings.Trim(port, ":")),
			Tags:    []string{ServiceTag},
			Address: "127.0.0.1",
		})
		if errr != nil {
			fmt.Printf("failed to listen(%v): %v", port, errr)
		}
		errr = grpcServer.Serve(lis)
		if errr != nil {
			fmt.Printf("failed to listen(%v): %v", port, errr)
			sw.Done()
		}
	}()
}

func startBadTasteGrpc(grpcServer *grpc.Server, port string, sw *sync.WaitGroup) {
	fmt.Printf("startBadTasteGrpc listen port(%v)\n", port)
	lis, _ := coreStartGrpc(grpcServer, port, sw)
	err := util.RegisterGrpcServiceToConsul(&util.RegisterGrpcServiceArgs{
		Name:    ServiceName,
		Port:    util.StrToInt(strings.Trim(port, ":")),
		Tags:    []string{ServiceTag},
		Address: "127.0.0.1",
	})
	if err != nil {
		fmt.Printf("failed RegisterGrpcServiceToConsul(%v): %v\n", port, err)
	}

	timeout := time.Second * time.Duration(rand.Intn(5)+5)
	maxTimes := 1000
	fmt.Printf("timeout: %v \n", timeout)
	retC := make(chan int, 1)
	go func() {
		i := 0
		for {
			select {
			case <-retC:
				fmt.Printf("after %v shutdown %v\n", timeout, port)
				err := lis.Close()
				if err != nil {
					fmt.Printf("lis close err %s\n", err)
				}
				sw.Done()
			case <-time.After(timeout * 2):
				fmt.Printf("[%v],after %v try shutdown port(%v) and wait restart\n", i, timeout, port)
				err := lis.Close()
				if err != nil {
					fmt.Printf("lis close err %s\n", err)
				}
				//util.ListenConsulStat(ServiceName, ServiceTag)
				time.Sleep(timeout)
				//util.ListenConsulStat(ServiceName, ServiceTag)
				lis, _ = coreStartGrpc(grpcServer, port, sw)
				i++
			}
			if i > maxTimes {
				break
			}
			if i == maxTimes {
				fmt.Printf("timeout max arraive: %v \n try stop", i)
				retC <- 1 //5次后强制退出并不重启
				i++
			}

		}
	}()
}

func coreStartGrpc(grpcServer *grpc.Server, port string, sw *sync.WaitGroup) (net.Listener, error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen port(%v): %v\n", port, err)
		sw.Done()
		return lis, err
	}
	go func() {
		errr := grpcServer.Serve(lis)
		if errr != nil {
			fmt.Printf("failed to Serve port(%v): %v\n", port, errr)
		}
		fmt.Printf("=======================================over=======================================\n")
	}()
	return lis, err
}
