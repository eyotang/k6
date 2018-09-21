package tcp

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/eyotang/load/library/binarypack"
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

func (t *TCP) Send(ctx context.Context, format []string, headers []interface{}, message []byte) (err error) {
	var (
		header []byte
		bp     = new(binarypack.BinaryPack)
	)
	if header, err = bp.Pack(format, headers); err != nil {
		log.Error("%+v", err)
		return
	}
	fmt.Printf("Send  ====> format: %v, headers: %v, message: %v\n", format, header, message)
	return
}

func (t *TCP) Close() (err error) {
	err = t.conn.Close()
	return
}
