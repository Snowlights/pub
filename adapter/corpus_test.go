package adapter

import (
	"context"
	"fmt"
	corpus "github.com/Snowlights/pub/grpc"
	"testing"
)

func TestLoginUser(t *testing.T) {
	ctx := context.Background()
	req := &corpus.LoginUserReq{
		EMail:                "858777157@qq.com",
		UserPassword:         "woaini12",
		Phone:                "",
		Code:                 0,
	}

	res := LoginUser(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}