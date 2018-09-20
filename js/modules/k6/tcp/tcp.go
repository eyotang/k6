package tcp

import (
	"context"
	"net"
	"strconv"

	"fmt"
	"github.com/pkg/errors"
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

func (t *TCP) Send(ctx context.Context, format []string, headers []interface{}, message []byte) {
	fmt.Printf("Send  ====> format: %v, headers: %v, message: %v\n", format, headers, message)
	return
}

func (t *TCP) Close() (err error) {
	err = t.conn.Close()
	return
}
