package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"coolcar/server/shared/auth/token"
	"coolcar/server/shared/id"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	authorizationPrefix = "Bearer "
)

// 实现 UnaryServerInterceptor 拦截器
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannnot open public key file: %v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key file: %v", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key: %v", err)
	}
	i := &interceptor{
		publicKey: pubKey,
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandleReq, nil
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	// 获取 accountID
	aid, err := i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}
	return handler(ContextWithAccountID(ctx, id.AccountID(aid)), req)
}

func tokenFromContext(ctx context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")

	// 获取请 求中的metadata
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "")
	}

	tkn := ""
	for _, v := range m[authorizationHeader] {
		// 判断 token 是否有 Bearer
		if strings.HasPrefix(v, authorizationPrefix) {
			tkn = v[len(authorizationPrefix):]
		}
	}

	if tkn == "" {
		return "", unauthenticated
	}

	return tkn, nil
}

type accountIDKey struct{}

// 用 accountId 生成 context 
func ContextWithAccountID(ctx context.Context, aid id.AccountID) context.Context {
	return context.WithValue(ctx, accountIDKey{}, aid)
}

// 从 context 中获取 accountID 
func AccountIDFromContext(ctx context.Context) (id.AccountID, error) {
	v := ctx.Value(accountIDKey{})
	// 判断获取的值是否是 string 类型
	aid, ok := v.(id.AccountID)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "")
	}
	return aid, nil
}

type interceptor struct {
	publicKey *rsa.PublicKey
	verifier  tokenVerifier
}

type tokenVerifier interface {
	Verify(token string) (string, error)
}
