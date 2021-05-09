/**
  @author: cheney
  @date: 2021/5/9
  @note:
 **/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaocanwei/sso-cloud/utils/midlog"
	"log"

	"net/http"
)

func setup() (e *gin.Engine, err error) {
	return
}

func main()  {
	/*ctx, cancel := context.WithCancel(context.Background())
	defer cancel()*/

	errorWriter := midlog.ErrorLog.Writer()
	defer errorWriter.Close()

	e, err := setup()
	if err != nil {
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", "8890"),
		Handler: e,
		ErrorLog: log.New(errorWriter, "", 0),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		}
	}()
}