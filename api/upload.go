package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go-gin-shop/enter/dto"
	"go-gin-shop/global"
)

type UploadApi struct {
}

// Save The detailed information:
// @Title Save
// @Description blog保存本地
// @Param c
func (UploadApi) Save(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, dto.Err("图片不存在"))
		return
	}
	file.Filename = uuid.NewV4().String() + ".png"
	// 文件保存到nginx目录中
	err = c.SaveUploadedFile(file, global.BlogFileSaveLocal+file.Filename)
	if err != nil {
		c.JSON(200, dto.Err("图片保存失败"))
		return
	}
	c.JSON(200, dto.OkData("/blogs/"+file.Filename))
}

// Delete The detailed information:
// @Title Delete
// @Description blog删除
// @Param c
func (UploadApi) Delete(c *gin.Context) {
	// TODO
}
