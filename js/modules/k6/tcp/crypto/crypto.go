package crypto

import (
	"context"
	"encoding/json"

	xcrypto "github.com/eyotang/load/library/crypto"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
)

type DesCrypto struct {
	Key     string `json:"key"`
	Mode    string `json:"mode"`
	IV      string `json:"iv"`
	PadMode string `json:"padmode"`
}

type Crypto struct {
}

func New() *Crypto {
	return &Crypto{}
}

func (c *Crypto) des(key string, mode string, iv string, padmode string) (des *xcrypto.Des, err error) {
	var (
		md      = uint8(xcrypto.CBC)
		padMode = uint8(xcrypto.PAD_PKCS5)
	)

	if mode == "ECB" {
		md = uint8(xcrypto.ECB)
	}
	if padmode == "PAD_NORMAL" {
		padMode = uint8(xcrypto.PAD_NORMAL)
	}

	if des, err = xcrypto.NewDes([]byte(key), md, []byte(iv), padMode); err != nil {
		err = errors.Wrapf(err, "crypto.NewDes failed")
		return
	}

	return
}

func (c *Crypto) Encrypt(ctx context.Context, dc string, input []byte) []byte {
	var (
		err error
		des *xcrypto.Des
		d   = &DesCrypto{}
	)

	if err = json.Unmarshal([]byte(dc), d); err != nil {
		log.Error("%+v", err)
		return nil
	}
	if des, err = c.des(d.Key, d.Mode, d.IV, d.PadMode); err != nil {
		log.Error("%+v", err)
		return nil
	}
	return des.Encrypt(input)
}

func (c *Crypto) Decrypt(ctx context.Context, dc string, input []byte) []byte {
	var (
		err error
		des *xcrypto.Des
		d   = &DesCrypto{}
	)

	if err = json.Unmarshal([]byte(dc), d); err != nil {
		log.Error("%+v", err)
		return nil
	}
	if des, err = c.des(d.Key, d.Mode, d.IV, d.PadMode); err != nil {
		log.Error("%+v", err)
		return nil
	}
	return des.Decrypt(input)
}
