package main

import (
 "fmt"
 "image"
 "image/jpeg"
 _ "image/png" // 需要引入这个包，才能解码 png 格式的图片
 "log"
 "os"
)

func main() {
 // 打开图片文件
 file, err := os.Open("original.jpg")
 if err != nil {
  log.Fatal(err)
 }
 defer file.Close()

 // 解码图片
 img, _, err := image.Decode(file)
 if err != nil {
  log.Fatal(err)
 }

 // 创建输出文件
 out, err := os.Create("compressed.jpg")
 if err != nil {
  log.Fatal(err)
 }
 defer out.Close()

 // 设置压缩质量
 options := &jpeg.Options{
  Quality: 50,
 }

 // 压缩图片并输出到文件
 err = jpeg.Encode(out, img, options)
 if err != nil {
  log.Fatal(err)
 }

 // 输出压缩后的图片文件大小
 info, err := out.Stat()
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println("压缩后的图片文件大小：", info.Size())
}
