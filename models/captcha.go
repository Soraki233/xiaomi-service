package models

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

//创建store
var store = base64Captcha.DefaultMemStore

// MakeCaptcha 生成验证码
func MakeCaptcha() (id string, b64s string, err error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuiopasdfghjklzxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
	}
	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()
	return id, b64s, err
}

// VerifyCaptcha 校验验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	fmt.Println(id)
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
