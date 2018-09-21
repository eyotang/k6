package tcp

import (
	"context"
	"fmt"
	"github.com/eyotang/load/library/header"
	"net"
	"strconv"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type TCP struct {
	ctx  context.Context
	conn net.Conn
}

func New() *TCP {
	return &TCP{}
}

func (t *TCP) Connect(ctx context.Context, host string, port uint64) (err error) {
	var (
		conn net.Conn
	)

	if conn, err = net.Dial("tcp", host+":"+strconv.FormatUint(port, 10)); err != nil {
		errors.Wrapf(err, "net.Dial failed")
		return
	}

	t.ctx = ctx
	t.conn = conn
	return
}

func (t *TCP) Pack(ctx context.Context, format []string, headers []interface{}, message []byte) (data []byte, err error) {
	var (
		hd []byte
		h  = header.New()
	)
	if hd, err = h.Pack(format, headers); err != nil {
		log.Error("%+v", err)
		return
	}
	data = append(hd, message...)
	return
}

func (t *TCP) Send(ctx context.Context, data []byte) (err error) {
	fmt.Printf("Send  ====> message: %v\n", data)
	return
}

func (t *TCP) Close() (err error) {
	err = t.conn.Close()
	return
}
