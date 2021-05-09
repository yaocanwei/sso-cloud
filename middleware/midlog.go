/**
  @author: cheney
  @date: 2021/5/9
  @note:
 **/
package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yaocanwei/sso-cloud/utils/midlog"
	"net/http"
)

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
	body bytes.Buffer
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)

	return
}

func AccessLogging (e *gin.Engine) gin.HandlerFunc {

	// 创建一个新的handler包装http.HandlerFunc
	return gin.HandlerFunc(func(c *gin.Context) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		logEntry := midlog.AccessLog.WithFields(logrus.Fields{
			"ip": c.Request.RemoteAddr,
			"method": c.Request.Method,
			"path": c.Request.RequestURI,
			"query": c.Request.URL.RawQuery,
			"request_body": buf.String(),

		})

		wc := &ResponseWithRecorder{
			ResponseWriter: c.Writer,
			statusCode: http.StatusOK,
			body: bytes.Buffer{},
		}

		// 调用下一个中间件或者最终的handler处理程序
		e.ServeHTTP(wc, c.Request)

		defer logEntry.WithFields(logrus.Fields{
			"status": wc.statusCode,
			"response_body": wc.body.String(),
		}).Info()

	})
}
