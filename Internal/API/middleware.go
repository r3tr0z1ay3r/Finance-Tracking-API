package Api

import (
	"fmt"
	"net/http"

	"github.com/r3tr0z1ay3r/Finance-Tracking-API.git/Config"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	Config.LoadEnv("Config/.env")
	api_key := Config.GetEnv("API_KEY")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.Header.Get("X-API-Key")
		fmt.Printf("Key received is %v and the api is %v\n", key, api_key)
		if key != api_key {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
