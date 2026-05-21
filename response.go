package wechatbot

import "encoding/json"

type Response struct {
	Ret  int             `json:"ret"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func (r *Response) IsSuccess() bool {
	return r.Ret == 200
}
