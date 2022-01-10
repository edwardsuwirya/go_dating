package httpreq

import (
	"fmt"
)

type LoginReq struct {
	UserName string
	Password string
}

func (p *LoginReq) String() string {
	return fmt.Sprintf("LoginReq => UserName: %s", p.UserName)
}
