package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/daved/faasfun/hellosvc"
)

func main() {
	svc := &service{
		h: hellosvc.Handle,
	}

	lambda.StartHandler(svc)
}

type service struct {
	h hellosvc.Handler
}

// Invoke ...
func (svc *service) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	b := bytes.NewBuffer(payload)
	fmt.Println(b.String())

	r, err := svc.h(ctx, b)
	if err != nil {
		return []byte{}, err
	}
	b.Reset()

	if _, err = b.ReadFrom(r); err != nil {
		return []byte{}, err
	}

	return b.Bytes(), nil
}
