/**
  @author: cheney
  @date: 2021/5/9
  @note:
 **/
package request

import "net/url"

type urlRequestParse interface {
	intoURLValues() url.Values
	intoURLPathParams(*url.URL)
}

type bodyer interface {
	intoBody() ([]byte, error)
}
