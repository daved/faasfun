package hellosvc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Handler ...
type Handler func(context.Context, io.Reader) (io.Reader, error)

// Handle ...
func Handle(ctx context.Context, r io.Reader) (io.Reader, error) {
	req := &HelloReq{}

	if err := json.NewDecoder(r).Decode(req); err != nil {
		return nil, err
	}

	resp := &HelloResp{
		Msg:  fmt.Sprintf("Hello, %s. Have a great day.", req.Name),
		Time: time.Now().String(),
	}

	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(resp); err != nil {
		return nil, err
	}

	return b, nil
}

// HelloReq ...
type HelloReq struct {
	Name string `json:"name"`
}

// HelloResp ...
type HelloResp struct {
	Msg  string `json:"msg"`
	Time string `json:"time"`
}
