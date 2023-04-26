package botstuff

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
	"image/png"
	"strings"
)

func generateQrCodeB64EncodedForUrl(data string) (string, error) {
	c, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return "", err
	}
	image := c.Image(512)
	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	var buf strings.Builder
	buf.WriteString("data:image/png;base64,")
	err = encoder.Encode(base64.NewEncoder(base64.StdEncoding, &buf), image)
	return buf.String(), err
}
