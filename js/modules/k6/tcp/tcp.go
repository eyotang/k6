package tcp

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/eyotang/k6/js/common"
	"github.com/labstack/gommon/log"
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

	fmt.Println("Connect ------->")

	if conn, err = net.Dial("tcp", host+":"+strconv.FormatUint(port, 10)); err != nil {
		errors.Wrapf(err, "net.Dial failed")
		return
	}

	t.ctx = ctx
	t.conn = conn
	return
}

func (t *TCP) Send(ctx context.Context, data []byte) (err error) {
	var (
		size int
	)

	state := common.GetState(ctx)

	// Check rate limit *after* we've prepared a request; no need to wait with that part.
	if rpsLimit := state.RPSLimit; rpsLimit != nil {
		if err = rpsLimit.Wait(ctx); err != nil {
			return
		}
	}

	fmt.Printf("Send  ====> message: %v\n", data)
	if size, err = t.conn.Write(data); err != nil {
		log.Error("%+v", err)
	}
	if size != len(data) {
		log.Errorf("send message failed! expected: %d, actual: %d", len(data), size)
	}
	return
}

func (t *TCP) Close() (err error) {
	err = t.conn.Close()
	return
}
