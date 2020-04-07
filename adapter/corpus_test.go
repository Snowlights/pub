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
		EMail:                "xueyeup@163.com",
		UserPassword:         "123456",
		Phone:                "",
		Code:                 "",
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
		Phone:                "18846082154",
	}
	res := SendMessage(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}

func TestUpdateUserInfo2(t *testing.T) {
	initLogic()
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
	initLogic()
	ctx := context.Background()
	req := &corpus.ListUserInfoReq{
		Offset:               0,
		Limit:                10,
		EMail:                "",
		UserName:             "",
		Phone:                "18846082154",
		Cookie:               "zhangwei",
	}
	res := ListUserInfo(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Println(res)
}

func TestUpdateUserInfo(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.UpdateUserInfoReq{
		UserId:               9,
		UserName:             "美丽",
		EMail:                "xueyeup@163.com",
		Password:             "123456",
		Description:          "好美丽雅",
		Cookie:               "Q99coO44kHBGc5f3",
	}
	res := UpdateUserInfo(ctx,req)
	if res.Errinfo != nil{
		fmt.Printf("%v",res.Errinfo.Msg)
	}
	fmt.Printf("%v",res)

}

func TestLogoutUserInfo(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.LogoutUserInfoReq{
		Cookie:               "Q99coO44kHBGc5f3",
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
	initLogic()
	ctx := context.Background()
	req := &corpus.AddAuthReq{
		AuthCode:             "SERVICE_TEXT_CODE",
		AuthDescription:      "测试权限",
		ServiceName:          "/service/test",
		Cookie:               "zhangwei",
	}
	res := AddAuth(ctx,req)
	fmt.Printf("%v\n",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestUpdateAuth(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.UpdateAuthReq{
		Id:                   8,
		AuthCode:             "SERVICE_TEST_CODE",
		AuthDescription:      "测试代码1",
		ServiceName:          "/test/test",
		Cookie:               "zhangwei",
	}
	res := UpdateAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestDelAuth(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.DelAuthReq{
		Id:                   8,
		AuthCode:             "",
		Cookie:               "zhangwei",
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
	initLogic()
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
	initLogic()
	ctx := context.Background()
	req := &corpus.AddAdminUserReq{
		UserId:               9,
		Cookie:               "zhangwei",
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
	initLogic()
	ctx := context.Background()
	req := &corpus.DelAdminUserReq{
		UserId:               9,
		Cookie:               "zhangwei",
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
	initLogic()
	ctx := context.Background()
	req := &corpus.ListAdminUserReq{
		Offset:               0,
		Limit:                10,
		Cookie:               "zhangwei",
	}
	res := ListAdminUser(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}

}

func TestAddUserAuth(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.AddUserAuthReq{
		UserId:               9,
		AuthCode:             "SERVICE_AGE_CODE",
		Cookie:               "zhangwei",
	}
	res := AddUserAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestDelUserAuth(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.DelUserAuthReq{
		UserId:               9,
		AuthCode:             "SERVICE_KEYWORD_CODE",
		Cookie:               "zhangwei",
	}
	res := DelUserAuth(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestListUserAuth(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.ListUserAuthReq{
		UserId:               9,
		Limit:                10,
		Offset:               0,
		Cookie:               "zhangwei",
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

func TestRecognizeAge(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.RecognizeAgeReq{
		Audio:                "C:\\Users\\华硕\\Desktop\\pr\\evaluation\\eng\\en_sentence.wav",
		Cookie:               "Q99coO44kHBGc5f3",
	}
	res := RecognizeAge(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestEvaluation(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.EvaluationReq{
		Audio:                "C:\\Users\\华硕\\Desktop\\pr\\evaluation\\aa\\In\\raz_in_p9_text.mp3",
		Text:                 "in the mud.",
		Cookie:               "zhangwei",
	}
	res := Evaluation(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestAddTransAudio(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.AddTransAudioReq{
		OriginAudio:          "http://xia2.kekenet.com/Sound/2015/11/Nov25_5700351F4Y.mp3",
		AudioType:            corpus.AudioType_ACC,
		Cookie:               "zhangwei",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := AddTransAudio(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestRecognizeImage(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.RecognizeImageReq{
		File:  "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=169785117,3207160551&fm=26&gp=0.jpg",
		Cookie:               "zhangwei",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := RecognizeImage(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestTreansAudioToText(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.TransAudioToTextReq{
		Audio:                "http://xia2.kekenet.com/Sound/2015/11/Nov25_5700351F4Y.mp3",
		Cookie:               "zhangwei",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := TransAudioToText(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}

func TestGetKeyWord(t *testing.T) {
	initLogic()
	ctx := context.Background()
	req := &corpus.GetKeyWordReq{
		Text:                 "新闻文本，是指新闻作品或新闻报道的存在形式。广义的新闻文本与广义的新闻足对应的，包括消息文本和通讯文本；狭义仅指消息文本，即纯粹的新闻文本。新闻文本有四个基本特征，文本结构的简单性。结构形式相对单一，结构要素。内在要求缺一不可，易于新闻传播主体、收受主体把握和理解，叙事结构比较简单，大多采用与新闻事实客观结构（时空结构、因果逻辑等）相一致的方式展开。",
		Cookie:               "zhangwei",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	res := GetKeyWord(ctx,req)
	fmt.Printf("%v",res)
	if res.Errinfo != nil{
		fmt.Printf("%v\n",res.Errinfo.Msg)
	}
}