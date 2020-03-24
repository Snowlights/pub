package adapter

import (
	"context"
	"fmt"
	corpus "github.com/Snowlights/pub/grpc"
	"testing"
)
//ONkSy6HRfzRX7w1n
func TestLoginUser(t *testing.T) {
	ctx := context.Background()
	req := &corpus.LoginUserReq{
		EMail:                "",
		UserPassword:         "",
		Phone:                "18846082154",
		Code:                 831950,
	}

	res := LoginUser(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}

func TestSendMessage(t *testing.T) {
	ctx := context.Background()
	req := &corpus.SendMessageReq{
		Phone:                "18846082154",
	}
	res := SendMessage(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}

func TestListUserInfo(t *testing.T) {
	ctx := context.Background()
	req := &corpus.ListUserInfoReq{
		Offset:               0,
		Limit:                10,
		EMail:                "",
		UserName:             "",
		Phone:                "",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := ListUserInfo(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}
func TestUpdateUserInfo(t *testing.T) {
	ctx := context.Background()
	req := &corpus.UpdateUserInfoReq{
		UserId:               4,
		UserName:             "花开雪夜",
		Phone:                "18846082154",
		EMail:                "858777157@qq.com",
		Description:          "我爱她她吗",
		Cookie:               "ONkSy6HRfzRX7w1n",
	}
	res := UpdateUserInfo(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Printf("%v",res)

}

func TestDelUserInfo(t *testing.T) {
	ctx := context.Background()
	req := &corpus.DelUserInfoReq{
		UserId:               0,
		Cookie:               "",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := DelUserInfo(ctx,req)
	fmt.Printf("%v",res)
}