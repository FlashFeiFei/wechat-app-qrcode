package main

import (
	"bytes"
	"github.com/guapo-organizations/wechat-app-qrcode/draw"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func compound(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Println(err)
		w.Write([]byte("参数解析失败"))
		return
	}
	app_qrcode_type := req.Form.Get("appQrcodeType")
	mask_type := req.Form.Get("maskType")
	dstImg_type := req.Form.Get("dstImgType")
	appQrcodeType, err := strconv.ParseInt(app_qrcode_type, 10, 64)
	if err != nil {
		w.Write([]byte("appQrcodeType解析失败"))
		return
	}
	strconv.ParseInt(mask_type, 10, 64)
	maskType, err := strconv.ParseInt(app_qrcode_type, 10, 64)
	if err != nil {
		w.Write([]byte("maskType解析失败"))
		return
	}
	strconv.ParseInt(dstImg_type, 10, 64)
	dstImgType, err := strconv.ParseInt(app_qrcode_type, 10, 64)
	if err != nil {
		w.Write([]byte("dstImgType解析失败"))
		return
	}
	app_qrcode, _, err := req.FormFile("appQrcode")
	if err != nil {
		log.Println(err)
		w.Write([]byte("appQrcode解析失败"))
		return
	}
	defer app_qrcode.Close()
	mask, _, err := req.FormFile("mask")
	if err != nil {
		log.Println(err)
		w.Write([]byte("mask解析失败"))
		return
	}
	defer mask.Close()


	//图片合成
	dst_img, err := draw.CompoundQrCode(app_qrcode, appQrcodeType, mask, maskType)
	if err != nil {
		log.Println(err)
		w.Write([]byte("图片绘制失败"))
		return
	}

	bodyBuf := &bytes.Buffer{}
	draw.ImageEncode(bodyBuf, dst_img, dstImgType)
	out_img, _ := ioutil.ReadAll(bodyBuf)
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(out_img)
}

func main() {
	http.HandleFunc("/compound", compound)
	http.ListenAndServe(":18083", nil)
}
