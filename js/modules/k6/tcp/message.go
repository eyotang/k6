package tcp

import (
	"context"

	"github.com/eyotang/load/library/header"
	"github.com/labstack/gommon/log"
)

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
