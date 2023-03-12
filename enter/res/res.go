package res

import "github.com/gin-gonic/gin"

type Response struct {
	Success  bool        `json:"success,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
}

func write(c *gin.Context, response Response) {
	c.JSON(200, response)
}

func Ok(c *gin.Context) {
	write(c, Response{
		Success: true,
	})
}

func OkData(c *gin.Context, data interface{}) {
	write(c, Response{
		Success: true,
		Data:    data,
	})
}

func Err(c *gin.Context, errStr string) {
	write(c, Response{
		Success:  false,
		ErrorMsg: errStr,
	})
}
