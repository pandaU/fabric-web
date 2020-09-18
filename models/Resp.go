package models

type Resp struct {
	Code int  `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:data`
}

func (resp *Resp) setAll(code int,msg string,data interface{}) *Resp {
    resp.Code=code
    resp.Msg=msg
    resp.Data=data
	return resp
}
func RespSuccess(data interface{}) *Resp {
	resp :=&Resp{}
	return resp.setAll(200,"SUCCESS",data)
}
func RespError(msg string)  *Resp {
	resp :=&Resp{}
	return resp.setAll(500,msg,nil)
}