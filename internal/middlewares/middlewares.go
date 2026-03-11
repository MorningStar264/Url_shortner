package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MorningStar264/Url_shortner/internal/helper"
)

// HTTP middleware setting a value on the request context
func JWT_Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		access_token := r.Header.Get("Authentication")
		fmt.Println(access_token)	
		jwt_token := strings.TrimPrefix(access_token, "Bearer ")
		fmt.Println(jwt_token)	
		err := helper.VerifyToken(jwt_token)
		if err != nil {
			fmt.Fprintf(w, "Invalid access token")
		}
		next.ServeHTTP(w, r)
	})
}
