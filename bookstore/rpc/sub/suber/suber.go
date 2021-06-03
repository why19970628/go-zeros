// Code generated by goctl. DO NOT EDIT!
// Source: sub.proto

//go:generate mockgen -destination ./suber_mock.go -package suber -source $GOFILE

package suber

import (
	"context"

	"bookstore/rpc/sub/sub"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	SubReq  = sub.SubReq
	SubResp = sub.SubResp

	Suber interface {
		Check(ctx context.Context, in *SubReq) (*SubResp, error)
	}

	defaultSuber struct {
		cli zrpc.Client
	}
)

func NewSuber(cli zrpc.Client) Suber {
	return &defaultSuber{
		cli: cli,
	}
}

func (m *defaultSuber) Check(ctx context.Context, in *SubReq) (*SubResp, error) {
	client := sub.NewSuberClient(m.cli.Conn())
	return client.Check(ctx, in)
}