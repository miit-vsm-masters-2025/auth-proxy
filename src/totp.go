package main

import (
	"bytes"
	"image/png"

	"github.com/pquerna/otp/totp"
)

func InitTOTP(issuer, accountName string) bytes.Buffer {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: accountName,
	})
	if err != nil {
		panic(err)
	}

	// Convert TOTP key into a QR code encoded as a PNG image.
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	png.Encode(&buf, img)
	key.Secret()
	return buf

}
