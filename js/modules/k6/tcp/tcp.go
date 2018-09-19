package tcp

import (
	"context"
	"net"
	"strconv"
	
	"github.com/pkg/errors"
)

type TCP struct {
	ctx           context.Context
	conn          *net.Conn
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
	
	t.conn = &conn
	return 
}
