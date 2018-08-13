package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(f)
}

func f(req helloReq) (helloResp, error) {
	msg := fmt.Sprintf("Hello from AWS, %s. Have a great day.", req.Name)

	resp := helloResp{
		Msg:  msg,
		Time: time.Now().String(),
	}

	return resp, nil
}

type helloReq struct {
	Name string `json:"name"`
}

type helloResp struct {
	Msg  string `json:"msg"`
	Time string `json:"time"`
}
