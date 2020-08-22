package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

const TCPAddr = "127.0.0.1:8080"

func main() {

	go serviceTcp()
	clientTcp()

}

func clientTcp() {
	sw := sync.WaitGroup{}
	sw.Add(1)
	sw.Add(1)
	//go clientTcpHandle(TCPAddr, "a")
	go func() {
		defer sw.Done()
		clientTcpHandle(TCPAddr, "a")
	}()
	go func() {
		defer sw.Done()
		clientTcpHandle(TCPAddr, "b")
	}()
	sw.Wait()
}
func clientTcpHandle(path string, name string) {
	conn, err := net.Dial("tcp", path)
	if err != nil {
		fmt.Println("tcp client", err)
		return
	}
	defer func() {
		conn.Close()
	}()
	i := 1000
	for {
		send := fmt.Sprintf("%s(%d)\n", name, i)
		n, err := conn.Write([]byte(send))
		if err != nil {
			fmt.Printf("client write (%v,%v)\n", n, err)
		}
		i++
		fmt.Printf("client write servcie %s", send)
		res, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			fmt.Printf("client get eof (%v,%v)\n", n, err)
			return
		}
		if err != nil {
			fmt.Println("tcp ReadAll", err)
			return
		}
		fmt.Printf("client get servcie %s \n", string(res))
		if i > 1010 {
			break
		}
	}

}
func serviceTcp() {
	fmt.Println("tcp 服务端")
	lis, err := net.Listen("tcp", TCPAddr)
	defer lis.Close()
	if err != nil {
		fmt.Println("tcp 服务端", err)
		return
	}
	serviceTcpMessage(lis)

}

func serviceTcpMessage(lis net.Listener) {
	defer func() {
		lis.Close()
	}()
	i := 0
	flag := false
	for {
		if flag {
			break
		}
		fmt.Println("lis start Accept ", i)
		con, err := lis.Accept()
		if err != nil {
			fmt.Println("lis.Accept", err)
			return
		}
		go func(con net.Conn) {
			err := serviceHandleConnect(con, i)
			if err == io.EOF {
				fmt.Println("service eof close")
				flag = true
			}
			con.Close()
		}(con)
		i++
	}

}

func serviceHandleConnect(con net.Conn, i int) error {
	r, err := bufio.NewReader(con).ReadString('\n')
	if err != nil {
		fmt.Println("read err ", err)
		return err
	}
	fmt.Println("service get ", string(r), " i:", i)
	resp := fmt.Sprintf("%+v,%d\n", con, i)
	n, err := con.Write([]byte(resp))
	fmt.Println("service set ", resp)
	if err != nil {
		fmt.Println("service set  err ", n, err)
		return err
	}
	return serviceHandleConnect(con, i+1)
}
