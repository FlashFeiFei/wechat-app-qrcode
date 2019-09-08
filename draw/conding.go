package draw

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

const (
	JPEG_TYPE = iota
	PNG_TYPE
)

//合成新的二维码
func CompoundQrCode(app_qrcode io.Reader, app_qrcode_type int64, mask io.Reader, mask_type int64) (out_img image.Image, err error) {
	//图片解码
	app_qrcode_img, err := imageDecode(app_qrcode, app_qrcode_type)
	if err != nil {
		return nil, err
	}
	mask_img, err := imageDecode(mask, mask_type)
	if err != nil {
		return nil, err
	}

	//组合图片
	out_img, err = DrawAppQRCode(app_qrcode_img, mask_img)
	if err != nil {
		return nil, err
	}

	return out_img, nil
}

//图片解码
func imageDecode(src_img io.Reader, image_type int64) (img image.Image, err error) {

	if image_type == JPEG_TYPE {
		img, err = jpeg.Decode(src_img)
		if err != nil {
			return nil, fmt.Errorf("解码失败")
		}
		return img, nil
	}

	if image_type == PNG_TYPE {
		img, err = png.Decode(src_img)
		if err != nil {
			return nil, err
		}
		return img, nil
	}

	return nil, fmt.Errorf("图片格式%d找不到", image_type)
}

//图片编码
func ImageEncode(out io.Writer, img image.Image, img_type int64) (err error) {

	if img_type == PNG_TYPE {
		err = png.Encode(out, img)
		return
	}
	err = jpeg.Encode(out, img, nil)
	return
}
