package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const privateKey = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCnrNMPRpe815w5
RgmPhJUULJeMceGdd7/7I/ZA9EaJLlJtg2fE272PGwgu77IkBHjw1BN26cLuCV0U
2L3NFbEdQaONQbknOYJ/3PiJhmg2tRMB7/WBCx3HOp7yW1KuE9yxCK4b7es+m6KN
E7pug5STDqL5tEaCX6EFI74mP0DfZgYmBt28Gv3iM1Pkh6T0LD8vih8mLMZN8L8p
o3smufDfG/2rzHTEZ78Q+cOgaZNrl7RW9Wj44RzrR3KUgyeK5NNRChBvQtPHtlnq
Qz+AFjaEEVGJcLM6MUFdlcHdy1Fk4RmKoYnyucPHmYp1HxtDkoP8Q0J/ygXyNV0/
P7i6nptRAgMBAAECggEBAKIQeS+am6769xSkjTkafL0zHIeyys7Yn8aty6acdFDD
ZQhUqker3FwlVJOJOjV13S9ozCdzaeWJR16O32UKQlZ0yxANJlizTV8oxVCniLLX
8bc9p51CkVWvY1H80r1OlVDHgwGbxHSPGV4iY1/N7hz1WLDhfgUlSQ0ervtox9sj
wLGR/r9XgWf1mVMjwRszTyGMHD290qN0JgaQhsCud+muT61XFLhIV0dSz3ul4wlc
2IMhTitF+oIshVXzmd/LlOP0WPd/xSToFc6fzw+8KOvoACAq5oaitWo4b2tcIFaS
5L9pZfNtMTbF0/T1/7O/cBfCyQr4DekutGCHY8B8XAECgYEA1kp1o5obRyCzHaS3
Sd5d0T44au6fqfVUttfcQN1Jrelk67r6Z4pLOxjsBYqDTtKZsNtGnKtHZq0p20Gp
dUBwpXVuHN7BLK0XG1kpvPtL7QkY1NCUn1iCgZHX8FRKycYk5OumQcoj5iR9zECm
K8Y1XlaLTQ28RECkoJ7YSK7eg0kCgYEAyE+gSs92wUxNtIAZux0YbTY0SpU0+Usy
ZlxKsq8hynXJdg5opJrsaH0WxI7co3YCmNqyixJ7XJZJiA2gBJfEuuOnmmlJFhxx
nSQvT4BplMeoMSm3YwXKsxHYOgRQkT5wcI3aeeWlj2EPJ7DodcFDZU+kChT/4gUf
ljhQZoI2T8kCgYEAoh5Q44XaLzSI7dtIo0TsuzmUWynOdzlYbr/eLOB9NmpFZKXf
fWe9xKb1ILgK4R4pEgjCYhKegQpuiSci+cbXsgWmWYcYpCELQzBwiD2h0mE4fQCT
//1pNndM49ARiJc0IFA7RriT8jAXT+h1Dtb0VzuoRZInpYc2RSIHRO2u/6kCgYAB
xQpSfuC6tnTdSmBv0cL8fAUcP5M3PJ3WX6xdRcTTqBS+kUQFaET8a1Z/KA/09b7y
IMSBDAnA+Kbvp8cpIzoeuJIrgBCgPGIYlFBCsIy+PsFSpd6z5kIzMM4rPQyFK/sM
U4SBnTDIQoBCxoJXP/zbcUeuux7DnW35AshbD31xWQKBgQCxrNrGeQLeHlbxsEQ9
+GxenN2P7DcXy/jm3eLwyX2xNw+DBrArKQjDFIKAeJgIlf8rS8E7NSp9Prgf0qeF
MFY/B0WDMHfP1vbqV6ns2bfLvxpDsDxQtStMTEywbcPaj2pgsHHOfvn1DuUl43w2
MmSBz/gNqeIewL1x7Pyz1PFReQ==
-----END PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := NewJWTTokenGenerator("coolcar/auth", key)
	g.nowFuc = func() time.Time {
		return time.Unix(1628428834, 0)
	}
	tkn, err := g.GenerateToken("6102b2376d1a998272252489", 2 * time.Hour)
	if err != nil {
		t.Errorf("cannot generate token:%v", err)
	}
	want := `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg0MzYwMzQsImlhdCI6MTYyODQyODgzNCwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjEwMmIyMzc2ZDFhOTk4MjcyMjUyNDg5In0.XqQGNlCeij_aIJmQubE_OQ4L8fBep2T-nIpcRZvGbWRIHlptm-96T4UZZhA7rcQx1k8L0o2GPu8fKuSk2SBLU35jr7UKdHk0QwNnmaz6qUHRScQ7pw6BIjepVcNz1K5MCr6iR8zilTAGHNAI5KdBzbEgud-tzJc6Gzhf2IU7tmDBGVBrAYm3M2XkRHX-mJrT-i9dSYwpTLqVIBTbhOQhwOZFrqaM2wPWCFW9dwwPgncpfMwMN6TTjO5JqoqQAOSmJ3Om4kKXC1m8CX8yLWL3AA_HWleSwZs5SxgPHO6c5DkFA_3yhTllxL5J38qADvEjEb3EchL2nzNV3RWBwzk4kg`
	if tkn != want {
		t.Errorf("wrong token generated. want:%q; got:%q", want, tkn)
	}
}
