# wechat-app-qrcode
微信小程序码中间的那个头像更换每个月有次数限制，所以做了个更换小程序码中间的那个头像的小功能



# wechat-app-qrcode
是一个golang的包

# 提供了网络传输

运行服务端
```
go run server.go
```

运行客户端
```
go run client.go
```


curl 参数说明

接下来发送一个http请求即可得到替换过后的小程序码


- 请求类型post
- 请求编码类型   Content-Type multipart/form-data
- 请求参数 appQrcode stream类型 小程序码
- 请求参数 mask stream类型 头像
- 请求参数 appQrcodeType string类型 小程序码图片类型 0 jpeg 1 png
- 请求参数 maskType string类型 头像图片类型 0 jpeg 1 png
- 请求参数 dstImgType string类型    输出的小程序码类型 0 jpeg 1 png


响应是一张图片


# demo 测试

我提供了一个服务端，接口地址 http://106.12.76.73:18083/compound

往接口地址发送一个multipart/form-data请求,结果返回处理完的图片