package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const PublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA543yyyDaQM4cCKt/TjkA
LFKwH8u8v64x2nBCauXMKYilgKb8R6NPLKwKG0Wu7QpE4aUjX/0fbDC7jSpps4G3
WmUBqjcv44TebZPpvKmH/aIfDQ35Bxcx/bnnIPxjUT2qMwjjo969Srg1UBxEOBCD
kiikFtluM19v8w5/US8Y6crooaPIDaZk8wBJq3HKXUpJ7O6mSYBwj64hrfhy50pC
FfjqweHTqarrlruOqS0wwU76pth3k6ehiKZFb4EY0gJ33K3ZFNoDJyIUqzuEroqa
N5CnjMb7pFjnefuk4ASuBgqljoFYwRyvr2G2OYQGGeWT5LRgiIyeP6yrKiSF+WT6
4QIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T)  {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PublicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name string
		tkn string
		now time.Time
		want string
		wantErr bool
	}{
		{
			name: "valid_token",
			tkn: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2MTAzYTllNzFjYzk2MzdmMjE2YzBmNmYiLCJpc3MiOiJjb29sY2FyL2F1dGgiLCJleHAiOjE2Mjg0Nzc1MjMsImlhdCI6MTYyODQ3NzUyM30.1A9gJ-TG9Nj9f-B__HhnDvi2NdkrVU_GxYjxsupuEhZQ5pZelIyBdnZHoUkeSJkqwo9MRyabd4odV43W3ZOUpO620DQEqwrkRLAGXW50Uekjof5IOKRSJPbqf2Dz1BsTeNeFvVm252Ji626vNht-xA4xGqgjqUbtkFM3DCqpm9llT-N-8ij5sX6ZM2wk7P-NlR4UzmhHbF1bgEcjscSfUfsj_q6NEg81x000ASDC-ZiKa4eBZqGRWU_rrT3lNo-udbAAWPfL87QEdYF35gfiekG0ByUmMNDmcEhHOogeOKfXHrl-KXr1WZtX7tHNr7CvXdYUX9paro3dLezzhNmmtQ",
			now: time.Unix(1628577523, 0),
			want: "6103a9e71cc9637f216c0f6f",
		},
		{
			name: "token_expired",
			tkn: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2MTAzYTllNzFjYzk2MzdmMjE2YzBmNmYiLCJpc3MiOiJjb29sY2FyL2F1dGgiLCJleHAiOjE2Mjg0Nzc1MjMsImlhdCI6MTYyODQ3NzUyM30.1A9gJ-TG9Nj9f-B__HhnDvi2NdkrVU_GxYjxsupuEhZQ5pZelIyBdnZHoUkeSJkqwo9MRyabd4odV43W3ZOUpO620DQEqwrkRLAGXW50Uekjof5IOKRSJPbqf2Dz1BsTeNeFvVm252Ji626vNht-xA4xGqgjqUbtkFM3DCqpm9llT-N-8ij5sX6ZM2wk7P-NlR4UzmhHbF1bgEcjscSfUfsj_q6NEg81x000ASDC-ZiKa4eBZqGRWU_rrT3lNo-udbAAWPfL87QEdYF35gfiekG0ByUmMNDmcEhHOogeOKfXHrl-KXr1WZtX7tHNr7CvXdYUX9paro3dLezzhNmmtQ",
			now: time.Unix(1629477523, 0),
			wantErr: true,
		},
		{
			name: "token_expired",
			tkn: "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2MTAzYTllNzFjYzk2MzdmMjE2YzBmNmYiLCJpc3MiOiJjb29sY2FyL2F1dGgiLCJleHAiOjE2Mjg0Nzc1MjMsImlhdCI6MTYyODQ3NzUyM30.1A9gJ-TG9Nj9f-B__HhnDvi2NdkrVU_GxYjxsupuEhZQ5pZelIyBdnZHoUkeSJkqwo9MRyabd4odV43W3ZOUpO620DQEqwrkRLAGXW50Uekjof5IOKRSJPbqf2Dz1BsTeNeFvVm252Ji626vNht-xA4xGqgjqUbtkFM3DCqpm9llT-N-8ij5sX6ZM2wk7P-NlR4UzmhHbF1bgEcjscSfUfsj_q6NEg81x000ASDC-ZiKa4eBZqGRWU_rrT3lNo-udbAAWPfL87QEdYF35gfiekG0ByUmMNDmcEhHOogeOKfXHrl-KXr1WZtX7tHNr7CvXdYUX9paro3dLezzhNmmtQ",
			now: time.Unix(1629477523, 0),
			wantErr: true,
		},
		{
			name: "bad_token",
			tkn: "bad_token",
			now: time.Unix(1629477523, 0),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}

			accountId, err := v.Verify(c.tkn)
			if err != nil {
				t.Errorf("verification failed:%v", err)
			}
		
			if c.want != accountId {
				t.Errorf("wrong account id. want:%q; got:%q", c.want, accountId)
			}
		})
	}


	

	
}