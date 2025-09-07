package middleware

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"
	"strings"
	"web-basic/src/types"
)

const VerifyMsg = "verified"

func AuthHandler(next types.HandleFunc) types.HandleFunc {
	ignore := []string{"/login", "public/index.html"}

	return func(ctx *types.Context) {
		//  URL Prefix가 ignore에 포함되면 auth 체크 X
		for _, s := range ignore {
			if strings.HasPrefix(ctx.Request.URL.Path, s) {
				next(ctx)
				return
			}
		}
		if v, err := ctx.Request.Cookie("X_AUTH"); err == http.ErrNoCookie {

			ctx.Redirect("/login")
			return
		} else if err != nil {
			ctx.RenderErr(http.StatusInternalServerError, err)
		} else if Verify(VerifyMsg, v.Value) {
			//  쿠키 인증되면 다음 핸들러
			next(ctx)
			return
		}
		ctx.Redirect("/login")
	}

}

func Verify(msg, sign string) bool {
	return hmac.Equal([]byte(sign), []byte(Sign(msg)))
}

func Sign(msg string) string {
	secretKey := []byte("GO-LANG")
	if len(secretKey) == 0 {
		return ""
	}
	mac := hmac.New(sha1.New, secretKey)
	io.WriteString(mac, msg)
	return hex.EncodeToString(mac.Sum(nil))
}
