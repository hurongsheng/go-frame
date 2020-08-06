package util

import (
	"context"
	"errors"
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	healthpbv1 "google.golang.org/grpc/health/grpc_health_v1"
	"time"
)

var ErrorServiceNotAvailable = errors.New("service not available")

func RegisterHealthCheck(srv *grpc.Server) {
	healthpbv1.RegisterHealthServer(srv, grpchealth.NewServer())
}

type RegisterGrpcServiceArgs struct {
	Id            string
	Name          string
	Port          int
	Tags          []string
	Address       string
	CheckTimeOut  string
	CheckInterval string
	CheckUrl      string
}

func RegisterGrpcServiceToConsul(args *RegisterGrpcServiceArgs) error {
	args.Id = fmt.Sprintf("%v_%v", args.Name, args.Port)
	args.CheckTimeOut = "1s"
	args.CheckInterval = "1s"
	//args.CheckInterval = fmt.Sprintf("%vs", rand.Intn(4)+2)
	args.CheckUrl = fmt.Sprintf("%v:%v", args.Address, args.Port)
	fmt.Printf("RegisterGrpcService: %+v \n", args)
	client, err := consulApi.NewClient(consulApi.DefaultConfig())
	if err != nil {
		return err
	}
	err = client.Agent().ServiceRegister(&consulApi.AgentServiceRegistration{
		ID:      args.Id,
		Name:    args.Name,
		Port:    args.Port,
		Tags:    args.Tags,
		Address: args.Address,
		Check: &consulApi.AgentServiceCheck{
			Interval: args.CheckInterval,
			Timeout:  args.CheckTimeOut,
			GRPC:     args.CheckUrl,
		},
	})
	return err
}

type ServiceStorage struct {
	services []*consulApi.ServiceEntry
	index    int
	name     string
	tag      string
}

func (s *ServiceStorage) GetService() *consulApi.ServiceEntry {
	ss := s.services[s.index/len(s.services)]
	s.index++
	return ss
}

func (s *ServiceStorage) GetSafeServiceConn(ctx context.Context) (*grpc.ClientConn, error) {
	if len(s.services) == 0 {
		return nil, ErrorServiceNotAvailable
	}
	ss := s.services[s.index%len(s.services)]
	s.index++
	fmt.Printf("try connect grpc %v:%v\n", ss.Service.Address, ss.Service.Port)
	ctxTimeOut, _ := context.WithTimeout(ctx, time.Second)
	return grpc.DialContext(ctxTimeOut, fmt.Sprintf("%v:%v", ss.Service.Address, ss.Service.Port), grpc.WithInsecure(), grpc.WithBlock())
}

func (s *ServiceStorage) listenPollingService() {
	for {
		select {
		case <-time.After(time.Second * 10):
			s.PollingService(true)
		}
	}
}
func (s *ServiceStorage) PollingService(polling bool) {
	fmt.Printf("PollingService ")
	ss, err := GetHealthService(s.name, s.tag, polling)
	if err != nil {
		fmt.Printf("PollingService err %s\n", err)
		return
	}
	fmt.Printf("now len(%v) \n", len(ss.services))
	s.services = ss.services
}

func GetHealthService(name string, tag string, polling bool) (*ServiceStorage, error) {
	client, err := consulApi.NewClient(consulApi.DefaultConfig())
	if err != nil {
		fmt.Printf("new client.Agent().Services() err:%s \n", err)
		return nil, err
	}
	services, _, err := client.Health().Service(name, tag, true, &consulApi.QueryOptions{})
	if err != nil {
		fmt.Printf("get client.Agent().Services() err:%s \n", err)
		return nil, err
	}
	ss := ServiceStorage{
		services: services,
		name:     name,
		tag:      tag,
	}
	if !polling {
		go ss.listenPollingService()
	}
	return &ss, nil
}
