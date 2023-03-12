package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"image"
	"image/jpeg"
	"image/png"
	"math/big"
	"os"
	"strconv"
)

type captchaModel struct {
	SliderImg string `json:"sliderImg,omitempty"`
	BgImg     string `json:"bgImg,omitempty"`
	X         int32  `json:"x,omitempty"`
	Y         int32  `json:"y,omitempty"`
}
type captchaVo struct {
	SliderImg string `json:"sliderImg,omitempty"`
	BgImg     string `json:"bgImg,omitempty"`
	Y         int32  `json:"y,omitempty"`
}

func CaptchaModelToVo(c *captchaModel) captchaVo {
	return captchaVo{
		SliderImg: c.SliderImg,
		BgImg:     c.BgImg,
		Y:         c.Y,
	}
}

const (
	bg    = "./static/code/bg/"
	mask  = "./static/code/mask.png"
	newBg = "./static/code/newBg/"
)

func GetRandInt(max int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(max-1)))
	return int(num.Int64())
}
func CreateCode() captchaModel {
	//生成随机数,用来随机选取图片
	nums := GetRandInt(10)
	//用于生成的图片名称
	imageId := uuid.New().String()
	//获取图片
	f, _ := os.Open(bg + strconv.Itoa(nums) + ".png")
	//获取随机x坐标
	imageRandX := GetRandInt(480 - 100)
	if imageRandX < 200 {
		imageRandX += 200
	}
	//获取随机y坐标
	imageRandY := GetRandInt(240 - 100)
	if imageRandY < 100 {
		imageRandY += 100
	}
	//转化为image对象
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	//设置截取的最大坐标值和最小坐标值
	maxPotion := image.Point{
		X: imageRandX,
		Y: imageRandY,
	}
	minPotion := image.Point{
		X: imageRandX - 100,
		Y: imageRandY - 100,
	}
	subimg := image.Rectangle{
		Max: maxPotion,
		Min: minPotion,
	}
	f, err = os.Create(newBg + imageId + "screenshot.jpeg")
	defer f.Close()
	//截取图像
	data := imaging.Crop(m, subimg)
	jpeg.Encode(f, data, nil)
	//base64编码
	buffer := bytes.NewBuffer(nil)
	jpeg.Encode(buffer, data, nil)
	maskBase64 := base64.StdEncoding.EncodeToString(buffer.Bytes())
	//设置遮罩
	bkBase64 := createCodeImg(bg+strconv.Itoa(nums)+".png", minPotion, imageId)
	captchaModel := &captchaModel{
		SliderImg: maskBase64,
		BgImg:     bkBase64,
		X:         int32(imageRandX),
		Y:         int32(imageRandY),
	}
	return captchaModel
}
func createCodeImg(path string, minPotion image.Point, imageId string) string {
	bg, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	maskFile, err := os.Open(mask)
	if err != nil {
		panic(err)
	}
	bgimg, err := png.Decode(bg)
	maskimg, err := png.Decode(maskFile)
	//参数：背景图，遮盖图，坐标，透明度
	data := imaging.Overlay(bgimg, maskimg, minPotion, 1.0)
	f, err := os.Create(newBg + imageId + ".jpeg")
	defer f.Close()
	jpeg.Encode(f, data, nil)
	//base64编码
	buffer := bytes.NewBuffer(nil)
	jpeg.Encode(buffer, data, nil)
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}
