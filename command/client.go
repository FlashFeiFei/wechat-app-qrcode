package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type MyMultipart struct {
	//这个不是io的writer，是multipart的结构体
	MultipartBodyWriter *multipart.Writer
}

func (mm MyMultipart) addField(fieldname, value string) (err error) {
	err = mm.MultipartBodyWriter.WriteField(fieldname, value)
	return
}

//部件方式，文件
func (mm MyMultipart) addFile(fieldname, filename, filepath string) (err error) {
	//打开文件
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer, err := mm.MultipartBodyWriter.CreateFormFile(fieldname, fieldname)
	if err != nil {
		return err
	}
	//文件写入
	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	client := http.Client{}

	bodyBuf := &bytes.Buffer{}
	mm := MyMultipart{MultipartBodyWriter: multipart.NewWriter(bodyBuf)}
	mm.addFile("appQrcode", "小程序码", "../resource/gh_1aa017b9d2a7_430.jpg")
	mm.addFile("mask", "替换图", "../resource/timg.jpg")
	mm.addField("appQrcodeType", "0")
	mm.addField("maskType", "0")
	mm.addField("dstImgType", "0")
	//必须要发之前关闭
	mm.MultipartBodyWriter.Close()

	request, err := http.NewRequest("post", "http://localhost:18083/compound", bodyBuf)

	request.Header.Set("Content-Type", mm.MultipartBodyWriter.FormDataContentType())

	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	result, _ := ioutil.ReadAll(response.Body)
	response.Body.Close() //释放资源
	f, _ := os.Create("./xcxm.jpeg")
	f.Write(result)
	f.Close()
}
