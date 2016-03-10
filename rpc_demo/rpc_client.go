package main

import (
	"fmt"
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

func main() {

	client, err := rpc.DialHTTP("tcp", URL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req := Req{LanLon: "31.01,109.456"}

	reply := Resp{}

	for i := 0; i < 10; i++ {
		e := client.Call("FsService.SearchVenues", &req, &reply)

		if e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println("rpc call result=", reply, reply.Venue.Id, reply.Venue.Name)
		}
	}

}
