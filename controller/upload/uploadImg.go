package upload

import (
	"flea-market/common/tools"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {

	f, err := c.FormFile("imgfile")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "上传失败!",
		})
		return
	} else {

		fmt.Println("file name is : ", f.Filename)

		fileExt := strings.ToLower(path.Ext(f.Filename))
		fmt.Println("file ext is : ", fileExt)
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			// c.JSON(200, gin.H{
			// 	"code": 400,
			// 	"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
			// })
			// return
		}
		fileName := tools.CreateUid()
		// fildDir := fmt.Sprintf("%s%d%s/", "static", time.Now().Year(), time.Now().Month().String())
		fildDir := fmt.Sprintf("%s/%d-%d-%d/", "static", time.Now().Year(), time.Now().Month(), time.Now().Day())

		os.Mkdir(fildDir, os.ModePerm)

		filepath := fmt.Sprintf("%s%s%s", fildDir, fileName, fileExt)
		c.SaveUploadedFile(f, filepath)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传成功!",
			"result": gin.H{
				"path": filepath,
			},
		})
	}
}
