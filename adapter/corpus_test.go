package adapter

import (
	"context"
	"fmt"
	corpus "github.com/Snowlights/pub/grpc"
	"testing"
)
//ONkSy6HRfzRX7w1n
func TestLoginUser(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.LoginUserReq{
		EMail:                "",
		UserPassword:         "",
		Phone:                "15535259636",
		Code:                 "343222",
	}

	res := LoginUser(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}

func TestSendMessage(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.SendMessageReq{
		Phone:                "15535259636",
	}
	res := SendMessage(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}

func TestUpdateUserInfo2(t *testing.T) {
	ctx := context.Background()
	req := &corpus.UpdateUserPhoneReq{
		Phone:                "1555555545",
		Code:                 "949116",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := UpdateUserPhone(ctx,req)
	fmt.Printf("%v",res)
}

func TestListUserInfo(t *testing.T) {
	ctx := context.Background()
	req := &corpus.ListUserInfoReq{
		Offset:               0,
		Limit:                10,
		EMail:                "",
		UserName:             "",
		Phone:                "18846082154",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := ListUserInfo(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}

func TestListUserInfo2(t *testing.T) {
	ctx := context.Background()
	req := &corpus.ListUserInfoReq{
		Offset:               0,
		Limit:                1,
		EMail:                "",
		UserName:             "",
		Phone:                "18846082154",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := ListUserInfo(ctx,req)
	fmt.Println(res)
}
func TestUpdateUserInfo(t *testing.T) {
	ctx := context.Background()
	req := &corpus.UpdateUserInfoReq{
		UserId:               6,
		UserName:             "美丽",
		EMail:                "858777157@qq.com",
		Password:             "123456",
		Description:          "好美丽雅",
		Cookie:               "tTawF21P99KlqUAE",
	}
	res := UpdateUserInfo(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Printf("%v",res)

}

func TestLogoutUserInfo(t *testing.T) {
	ctx := context.Background()
	req := &corpus.LogoutUserInfoReq{
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := LogoutUserInfo(ctx,req)
	fmt.Printf("%v",res)
}

func TestDelUserInfo(t *testing.T) {
	ctx := context.Background()
	req := &corpus.DelUserInfoReq{
		UserId:               6,
		Cookie:               "tTawF21P99KlqUAE",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := DelUserInfo(ctx,req)
	fmt.Printf("%v",res)
}

func TestAddAuth(t *testing.T) {
	ctx := context.Background()
	req := &corpus.AddAuthReq{
		AuthCode:             "SERVICE_AGE_CODE",
		AuthDescription:      "年龄识别服务（主要用于内部，确定发展方向）",
		ServiceName:          "/service/age",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := AddAuth(ctx,req)
	fmt.Printf("%v\n",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestUpdateAuth(t *testing.T) {
	ctx := context.Background()
	req := &corpus.UpdateAuthReq{
		Id:                   6,
		AuthCode:             "SERVICE_AUDIO_RECOGNIZE_CODE",
		AuthDescription:      "音频识别服务",
		ServiceName:          "/audio",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := UpdateAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestDelAuth(t *testing.T) {
	ctx := context.Background()
	req := &corpus.DelAuthReq{
		Id:                   0,
		AuthCode:             "SERVICE_AUDIO_RECOGNIZE_CODE",
		Cookie:               "JnL3gxsI402j4hs4",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := DelAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestListAuth(t *testing.T) {
	ctx := context.Background()
	req := &corpus.ListAuthReq{
		Offset:               0,
		Limit:                10,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := ListAuth(ctx,req)
	fmt.Printf("%v",res)
}

func TestAddAdminUser(t *testing.T) {
	ctx := context.Background()
	req := &corpus.AddAdminUserReq{
		UserId:               8,
		Cookie:               "JnL3gxsI402j4hs4",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := AddAdminUser(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestDelAdminUser(t *testing.T) {
	ctx := context.Background()
	req := &corpus.DelAdminUserReq{
		UserId:               8,
		Cookie:               "JnL3gxsI402j4hs4",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := DelAdminUser(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestListAdminUser(t *testing.T) {
	ctx := context.Background()
	req := &corpus.ListAdminUserReq{
		Offset:               0,
		Limit:                10,
		Cookie:               "",
	}
	res := ListAdminUser(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}

}

func TestAddUserAuth(t *testing.T) {
	ctx := context.Background()
	req := &corpus.AddUserAuthReq{
		UserId:               4,
		AuthCode:             "SERVICE_AUDIO_CODE",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := AddUserAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestDelUserAuth(t *testing.T) {
	ctx := context.Background()
	req := &corpus.DelUserAuthReq{
		UserId:               4,
		AuthCode:             "SERVICE_AUDIO_CODE",
		Cookie:               "JnL3gxsI402j4hs4",
	}
	res := DelUserAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestListUserAuth(t *testing.T) {
	ctx := context.Background()
	req := &corpus.ListUserAuthReq{
		UserId:               4,
		Limit:                10,
		Offset:               0,
		Cookie:               "",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := ListUserAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

