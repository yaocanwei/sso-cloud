/**
  @author: cheney
  @date: 2021/5/9
  @note:
 **/
package store

import (
	"context"
	"github.com/gin-gonic/gin"
)

type StoreResp struct {
	username string `json:"username"`
	serviceTicket string `json:"st"`
}

const (
	key int = iota
	authKey
)

func setStoreResp(ctx *gin.Context, resp *StoreResp) {
	r := ctx.Request
	r2 := r.WithContext(context.WithValue(r.Context(), key, resp))
	*r = *r2
}

func getStore(ctx *gin.Context) *StoreResp {
	if rst := ctx.Value(authKey); rst != nil {
		return rst.(*StoreResp)
	}
	return nil
}

func IsAuth(ctx *gin.Context) bool {
	if rst := getStore(ctx); rst != nil {
		return true
	}
	return false
}

