package req

type TPageReq struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}
