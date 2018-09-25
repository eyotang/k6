package proto

import (
	"context"
	"reflect"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	xproto "github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type Proto struct {
}

func New() *Proto {
	return &Proto{}
}

func (p *Proto) NewMessage(ctx context.Context, json string, m string) (data []byte, err error) {
	var (
		r   = strings.NewReader(json)
		mt  = xproto.MessageType(m)
		v   = reflect.New(mt.Elem())
		msg = v.Interface().(xproto.Message)
	)
	if err = jsonpb.Unmarshal(r, msg); err != nil {
		err = errors.Wrapf(err, "proto.Unmarshal failed")
		return
	}

	if data, err = xproto.Marshal(msg); err != nil {
		err = errors.Wrapf(err, "proto.Marshal failed")
		return
	}
	return
}
