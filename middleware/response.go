/**
  @author: cheney
  @date: 2021/5/9
  @note:
 **/
package middleware

import "github.com/gin-gonic/gin"

func JsonResponse(c *gin.Context, code int, data interface{}, msg string)  {
	c.JSON(code, gin.H{"data": data, "msg": msg})
}
