package conn

import (
	"fmt"
	"net"
	"time"
)

type Status string

const (
	CLOSED Status = "CLOSED"
	OPEN          = "OPEN"
)

type Result struct {
	Status Status
	Err    error
}

func Check(protocol, address, port string) Result {
	conn, err := net.DialTimeout(protocol, fmt.Sprintf("%s:%s", address, port), 1*time.Second)
	if err != nil {
		return Result{
			Status: CLOSED,
			Err:    err,
		}
	}
	defer conn.Close()
	return Result{
		Status: OPEN,
		Err:    nil,
	}
}
