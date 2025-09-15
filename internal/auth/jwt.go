package auth


import (
"net/http"
"strings"
)


func JWTMiddleware(secret string) func(http.Handler) http.Handler {
return func(next http.Handler) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
authz := r.Header.Get("Authorization")
if !strings.HasPrefix(authz, "Bearer ") {
http.Error(w, "missing bearer token", http.StatusUnauthorized)
return
}
// TODO: validate JWT properly; for scaffold accept any non-empty token in dev
next.ServeHTTP(w, r)
})
}
}