package adapter

import (
	"context"
	corpus "github.com/Snowlights/pub/grpc"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"strconv"
)

const (
	address     = "localhost:50051"
	defaultName = "corpus"
)
var client corpus.CorpusServiceClient

func initLogic(){
	ctx := context.Background()
	Prepare(ctx)
}

func initConnect(ctx context.Context) *grpc.ClientConn{
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func Prepare(ctx context.Context){
	conn := initConnect(ctx)
	client = corpus.NewCorpusServiceClient(conn)
}

func getHashKey(ctx context.Context) string {
	routeKey := ""
	if ctx.Value("routeKey") != nil {
		val, ok := ctx.Value("routeKey").(string)
		if ok {
			routeKey = val
		}
	}
	if len(routeKey) < 1 {
		routeKey = strconv.Itoa(rand.Int())
	}
	return routeKey
}

func LoginUser(ctx context.Context,req *corpus.LoginUserReq) (res *corpus.LoginUserRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.LoginUser(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil {
		res = &corpus.LoginUserRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}

	return
}

func LogoutUserInfo(ctx context.Context,req *corpus.LogoutUserInfoReq) (res *corpus.LogoutUserInfoRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.LogoutUserInfo(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.LogoutUserInfoRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func UpdateUserPhone(ctx context.Context,req *corpus.UpdateUserPhoneReq) (res *corpus.UpdateUserPhoneRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.UpdateUserPhone(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.UpdateUserPhoneRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}


func SendMessage(ctx context.Context,req *corpus.SendMessageReq) (res *corpus.SendMessageRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.SendMessage(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.SendMessageRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func UpdateUserInfo(ctx context.Context,req *corpus.UpdateUserInfoReq)(res *corpus.UpdateUserInfoRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.UpdateUserInfo(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.UpdateUserInfoRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func DelUserInfo(ctx context.Context,req *corpus.DelUserInfoReq)(res *corpus.DelUserInfoRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.DelUserInfo(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.DelUserInfoRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func ListUserInfo(ctx context.Context,req *corpus.ListUserInfoReq) (res *corpus.ListUserInfoRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)

	res, err := client.ListUserInfo(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.ListUserInfoRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}

	return
}

func AddAdminUser(ctx context.Context,req *corpus.AddAdminUserReq)(res *corpus.AddAdminUserRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.AddAdminUser(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.AddAdminUserRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func DelAdminUser(ctx context.Context,req *corpus.DelAdminUserReq)(res *corpus.DelAdminUserRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.DelAdminUser(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.DelAdminUserRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}

	return
}

func ListAdminUser(ctx context.Context,req *corpus.ListAdminUserReq)(res *corpus.ListAdminUserRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.ListAdminUser(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.ListAdminUserRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func AddAuth(ctx context.Context,req *corpus.AddAuthReq)(res *corpus.AddAuthRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.AddAuth(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.AddAuthRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func UpdateAuth(ctx context.Context,req *corpus.UpdateAuthReq)(res *corpus.UpdateAuthRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.UpdateAuth(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.UpdateAuthRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func DelAuth(ctx context.Context,req *corpus.DelAuthReq) (res *corpus.DelAuthRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.DelAuth(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err!= nil{
		res = &corpus.DelAuthRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func ListAuth(ctx context.Context,req *corpus.ListAuthReq)(res *corpus.ListAuthRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.ListAuth(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.ListAuthRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func AddUserAuth(ctx context.Context,req *corpus.AddUserAuthReq) (res *corpus.AddUserAuthRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.AddUserAuth(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.AddUserAuthRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func DelUserAuth(ctx context.Context,req *corpus.DelUserAuthReq)(res *corpus.DelUserAuthRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.DelUserAuth(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.DelUserAuthRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func ListUserAuth(ctx context.Context,req *corpus.ListUserAuthReq)(res *corpus.ListUserAuthRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.ListUserAuth(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.ListUserAuthRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func AddTransAudio(ctx context.Context,req *corpus.AddTransAudioReq)(res *corpus.AddTransAudioRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.AddTransAudio(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.AddTransAudioRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func DelTransAudio(ctx context.Context,req *corpus.DelTransAudioReq)(res *corpus.DelTransAudioRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.DelTransAudio(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.DelTransAudioRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func ListTransAudio(ctx context.Context,req *corpus.ListTransAudioReq)(res* corpus.ListTransAudioRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.ListTransAudio(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.ListTransAudioRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func GetKeyWord(ctx context.Context,req *corpus.GetKeyWordReq)(res *corpus.GetKeyWordRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.GetKeyWord(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.GetKeyWordRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func Evaluation(ctx context.Context,req *corpus.EvaluationReq)(res *corpus.EvaluationRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	conn := initConnect(ctx)
	c := corpus.NewCorpusServiceClient(conn)
	res, err := c.Evaluation(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.EvaluationRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func RecognizeImage(ctx context.Context,req *corpus.RecognizeImageReq)(res *corpus.RecognizeImageRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.RecognizeImage(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.RecognizeImageRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func TreansAudioToText(ctx context.Context,req *corpus.TransAudioToTextReq)(res *corpus.TransAudioToTextRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.TransAudioToText(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.TransAudioToTextRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}

func RecognizeAge(ctx context.Context,req *corpus.RecognizeAgeReq)(res *corpus.RecognizeAgeRes){
	routeKey := getHashKey(ctx)
	log.Printf("%v routekey ",routeKey)
	res, err := client.RecognizeAge(ctx,req)
	log.Printf("%v success %v",ctx,res)
	if err != nil{
		res = &corpus.RecognizeAgeRes{
			Errinfo:              &corpus.ErrorInfo{
				Ret:                  -1,
				Msg:                  err.Error(),
				XXX_NoUnkeyedLiteral: struct{}{},
				XXX_unrecognized:     nil,
				XXX_sizecache:        0,
			},
			Data:                 nil,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
	}
	return
}