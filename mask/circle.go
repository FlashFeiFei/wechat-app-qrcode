package mask

import (
	"image"
	"image/color"
	"math"
)

//创建一个圆形覆盖物
func NewCircleMask(d int) CircleMask {
	return CircleMask{d}
}

//圆形覆盖物
type CircleMask struct {
	diameter int //直径
}

//图像的模型
func (ci CircleMask) ColorModel() color.Model {
	return color.RGBAModel
}

//图像的绘制区域，即面积,二维坐标
func (ci CircleMask) Bounds() image.Rectangle {
	return image.Rect(0, 0, ci.diameter, ci.diameter)
}

// At方法返回(x, y)位置的色彩
func (ci CircleMask) At(x, y int) color.Color {
	d := ci.diameter

	//绘制面积范围内的任意一点到原点的距离小于半径的，即
	dis := math.Sqrt(math.Pow(float64(x-d/2), 2) + math.Pow(float64(y-d/2), 2))
	if dis > float64(d)/2 {
		//点到原点的距离大于半径，原的范围外，完全透明的白色，代替原来的颜色
		return color.RGBA{255, 255, 255, 0}
	} else {
		//原内，用彩色笔
		return color.RGBA{0, 0, 255, 255}
	}
}
