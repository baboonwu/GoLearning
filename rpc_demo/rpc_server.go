package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

const (
	URL = "127.0.0.1:3003"
)

type Req struct {
	LanLon string
}

type Venue struct {
	Id   string
	Name string
}

type Resp struct {
	Venue Venue
}

// service
type FsService struct {
}

// Method 
func (fs *FsService) SearchVenues(req *Req, reply *Resp) error {

	fmt.Println("recv request=", req)

	*reply = Resp{
		Venue: Venue{
			Id:   "123456",
			Name: "test",
		},
	}

	fmt.Println("---- response=", reply)

	return nil
}

func main() {

	fs := new(FsService)
	rpc.Register(fs)
	rpc.HandleHTTP()

	err := http.ListenAndServe(URL, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
