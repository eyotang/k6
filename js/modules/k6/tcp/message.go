package tcp

import (
	"context"

	"github.com/eyotang/load/library/binarypack"
	"github.com/labstack/gommon/log"
)

func (t *TCP) Pack(ctx context.Context, format []string, headers []interface{}, message []byte) (data []byte, err error) {
	var (
		hd []byte
		bp = &binarypack.BinaryPack{}
	)
	if hd, err = bp.Pack(format, headers); err != nil {
		log.Error("%+v", err)
		return
	}
	data = append(hd, message...)
	return
}
