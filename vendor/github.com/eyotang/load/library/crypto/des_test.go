package crypto

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDesEncryptDecrypt(t *testing.T) {
	type params struct {
		mode         uint8
		padmode      uint8
		plaintext    []byte
		expectCipher []byte
	}

	var (
		targetPlain []byte
		ciphertext  []byte
		des         *Des
		pars        []params
		p           params
		err         error
	)

	Convey("TEST Des", t, func() {
		pars = []params{
			{
				mode:         CBC,
				padmode:      PAD_PKCS5,
				plaintext:    []byte{10, 5, 116, 101, 115, 116, 49, 18, 1, 49, 26, 1, 49, 32, 162, 224, 194, 181, 5, 42, 1, 49, 50, 1, 49, 58, 1, 49, 66, 10, 50, 48, 49, 56, 46, 48, 55, 46, 51, 49, 74, 1, 49, 80, 0, 88, 1, 98, 3, 52, 46, 48, 106, 0, 114, 0, 122, 0, 130, 1, 0},
				expectCipher: []byte{235, 181, 105, 171, 207, 26, 179, 80, 34, 70, 122, 11, 138, 9, 131, 254, 128, 10, 4, 147, 150, 232, 252, 175, 36, 201, 239, 51, 232, 60, 160, 171, 144, 219, 252, 248, 133, 245, 196, 120, 193, 19, 54, 112, 53, 150, 82, 242, 246, 35, 197, 14, 105, 217, 108, 114, 29, 135, 238, 207, 142, 235, 113, 145},
			},
			{
				mode:         ECB,
				padmode:      PAD_PKCS5,
				plaintext:    []byte{10, 5, 116, 101, 115, 116, 49, 18, 1, 49, 26, 1, 49, 32, 162, 224, 194, 181, 5, 42, 1, 49, 50, 1, 49, 58, 1, 49, 66, 10, 50, 48, 49, 56, 46, 48, 55, 46, 51, 49, 74, 1, 49, 80, 0, 88, 1, 98, 3, 52, 46, 48, 106, 0, 114, 0, 122, 0, 130, 1, 0},
				expectCipher: []byte{97, 157, 185, 35, 19, 138, 62, 174, 50, 70, 5, 255, 244, 126, 39, 249, 201, 50, 97, 117, 99, 159, 59, 116, 175, 119, 30, 44, 137, 83, 53, 56, 16, 38, 49, 92, 116, 61, 246, 16, 240, 48, 135, 90, 54, 110, 51, 139, 72, 135, 251, 129, 112, 193, 21, 81, 236, 194, 220, 190, 61, 190, 97, 68},
			},
			{
				mode:         CBC,
				padmode:      PAD_NORMAL,
				plaintext:    []byte{10, 5, 116, 101, 115, 116, 49, 18, 1, 49, 26, 1, 49, 32, 162, 224, 194, 181, 5, 42, 1, 49, 50, 1, 49, 58, 1, 49, 66, 10, 50, 48, 49, 56, 46, 48, 55, 46, 51, 49, 74, 1, 49, 80, 0, 88, 1, 98, 3, 52, 46, 48, 106, 0, 114, 0, 122, 0, 130, 1},
				expectCipher: []byte{235, 181, 105, 171, 207, 26, 179, 80, 34, 70, 122, 11, 138, 9, 131, 254, 128, 10, 4, 147, 150, 232, 252, 175, 36, 201, 239, 51, 232, 60, 160, 171, 144, 219, 252, 248, 133, 245, 196, 120, 193, 19, 54, 112, 53, 150, 82, 242, 246, 35, 197, 14, 105, 217, 108, 114, 135, 70, 122, 38, 205, 98, 62, 3},
			},
			{
				mode:         ECB,
				padmode:      PAD_NORMAL,
				plaintext:    []byte{10, 5, 116, 101, 115, 116, 49, 18, 1, 49, 26, 1, 49, 32, 162, 224, 194, 181, 5, 42, 1, 49, 50, 1, 49, 58, 1, 49, 66, 10, 50, 48, 49, 56, 46, 48, 55, 46, 51, 49, 74, 1, 49, 80, 0, 88, 1, 98, 3, 52, 46, 48, 106, 0, 114, 0, 122, 0, 130, 1},
				expectCipher: []byte{97, 157, 185, 35, 19, 138, 62, 174, 50, 70, 5, 255, 244, 126, 39, 249, 201, 50, 97, 117, 99, 159, 59, 116, 175, 119, 30, 44, 137, 83, 53, 56, 16, 38, 49, 92, 116, 61, 246, 16, 240, 48, 135, 90, 54, 110, 51, 139, 72, 135, 251, 129, 112, 193, 21, 81, 63, 62, 208, 153, 35, 142, 121, 97},
			},
		}
		for _, p = range pars {
			des, err = NewDes([]byte("TANGTANG"), p.mode, []byte("TANGTANG"), p.padmode)
			So(err, ShouldBeNil)

			ciphertext = des.Encrypt(p.plaintext)
			//fmt.Printf("===>%v\n", ciphertext)
			So(ciphertext, ShouldResemble, p.expectCipher)

			targetPlain = des.Decrypt(ciphertext)
			So(targetPlain, ShouldResemble, p.plaintext)
		}
	})
}
