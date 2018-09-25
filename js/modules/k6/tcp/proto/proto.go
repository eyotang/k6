package proto

import (
	"context"
	"io"

	"github.com/golang/protobuf/jsonpb"
	xproto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Proto struct {
}

func New() *Proto {
	return &Proto{}
}

func (p *Proto) NewMessage(ctx context.Context, r io.Reader, pb xproto.Message) (data []byte, err error) {
	pb = new(xproto.Message)
	if err = jsonpb.Unmarshal(r, pb); err != nil {
		err = errors.Wrapf(err, "proto.Unmarshal failed")
		return
	}

	if data, err = xproto.Marshal(pb); err != nil {
		err = errors.Wrapf(err, "proto.Marshal failed")
		return
	}
	return
}
