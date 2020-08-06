package main

import (
	"context"
	"encoding/json"
	"fmt"
	api "frame/api"
	"frame/util"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	logInfo("try http \n")
	ss, err := util.GetHealthService(api.ServiceName, api.ServiceTag, false)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	gate := Gate{
		ssClient: ss,
		mux:      mux,
	}
	gate.router("/vote", gate.List)
	gate.router("/rank", gate.List)

	logInfo("start http")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		logErr(err)
	}
}
func logInfo(message string) {
	fmt.Printf("message %v", message)
}

func logErr(err error) {
	logInfo(err.Error())
}

type Gate struct {
	ssClient *util.ServiceStorage
	mux      *http.ServeMux
}

func (g *Gate) router(path string, funcA func(ctx context.Context, request *http.Request, req url.Values) (interface{}, error)) {
	g.mux.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		ctxBg := context.Background()
		ctx, _ := context.WithTimeout(ctxBg, time.Second*30)
		res := make([]byte, 0)
		var err error
		defer func() {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			_, _ = writer.Write(res)
		}()
		req, err := g.getRequest(request)

		resInterface, err := funcA(ctx, request, req)
		if err != nil {
			return
		}
		res, err = json.Marshal(resInterface)
		if err != nil {
			return
		}
	})
}

func (g *Gate) getRequest(request *http.Request) (req url.Values, err error) {
	defer func() {
		fmt.Printf("req: %v \n", req)
	}()
	err = request.ParseForm()
	if err != nil {
		return nil, err
	}
	if request.Method != "GET" {
		return request.PostForm, nil
	} else {
		return request.Form, nil
	}
}

func (g *Gate) SServiceClient(ctx context.Context, f func(ctx context.Context, ss api.SServiceClient) (interface{}, error)) (interface{}, error) {
	conn, err := g.ssClient.GetSafeServiceConn(ctx)
	if err != nil {
		if err != context.DeadlineExceeded {
			return nil, err
		}
		fmt.Printf("GetSafeServiceConn time out service may down retry 1 tims\n")
		g.ssClient.PollingService(true)
		conn, err = g.ssClient.GetSafeServiceConn(ctx)
		if err != nil {
			return nil, err
		}
	}
	defer conn.Close()
	ssClient := api.NewSServiceClient(conn)
	return f(ctx, ssClient)
}
