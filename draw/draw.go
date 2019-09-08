package draw

import (
	"fmt"
	qrcode_mask "github.com/guapo-organizations/wechat-app-qrcode/mask"
	"image"
	"image/draw"
	"strconv"
)

//  小程序码/覆盖物 的比例
var ratio float64

func init() {
	ratio, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", 430./197.0), 64)
}

//src 为小程序码
//mask 为图片的图片
//结果是绘制出新图
func DrawAppQRCode(src image.Image, mask image.Image) (image.Image, error) {

	//通过小程序码获取计算覆盖物所需要的宽度
	calculate_mask_witdh, err := CalculatetMaskSizeBySrc(src)

	if err != nil {
		return nil, err
	}

	//获取覆盖物所需要的宽度
	mask_witdh, err := GetSuareBySize(mask)

	if err != nil {
		return nil, fmt.Errorf("覆盖物的图像不是正方形的")
	}

	//获取小程序码的宽度
	src_witdh, _ := GetSuareBySize(src)
	if mask_witdh != calculate_mask_witdh {
		return nil, fmt.Errorf("小程序码%d*%d,覆盖物大小必须要%d*%d", src_witdh, src_witdh, calculate_mask_witdh, calculate_mask_witdh)
	}

	//创建画布
	dstImg_canvas := image.NewRGBA(image.Rect(0, 0, src_witdh, src_witdh))
	//把小程序码画上去
	draw.Draw(dstImg_canvas, src.Bounds(), src, image.Pt(0, 0), draw.Over)

	//圆形覆盖物
	circle_mask := qrcode_mask.NewCircleMask(calculate_mask_witdh)

	//把圆形覆盖物绘制到图片
	//平移到画布的中心点
	move_pt := ((src_witdh - calculate_mask_witdh) / 2)

	draw.DrawMask(dstImg_canvas, src.Bounds().Add(image.Pt(move_pt, move_pt)), mask, image.Pt(0, 0), circle_mask, image.Pt(0, 0), draw.Over)

	return dstImg_canvas, nil

}

//通过图片来判断是否为方形
func ValidateSquareByImage(img image.Image) bool {
	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y

	if width == height {
		return true
	}

	return false
}

//获取图片的实际宽高
func GetSuareBySize(img image.Image) (int, error) {
	if !ValidateSquareByImage(img) {
		return 0, fmt.Errorf("图像不是方形")
	}

	width := img.Bounds().Max.X - img.Bounds().Min.X

	return width, nil
}

//通过原图，计算覆盖物的大小
func CalculatetMaskSizeBySrc(src image.Image) (int, error) {

	if !ValidateSquareByImage(src) {
		return 0, fmt.Errorf("小程序码不是正方形")
	}

	width := src.Bounds().Max.X - src.Bounds().Min.X

	return int(float64(width) / ratio), nil
}
