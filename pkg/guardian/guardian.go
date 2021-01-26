package guardian

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
)

// Guardian middleware
type Guardian struct {
	whitelist        map[string]string
	authenticateFunc basic.AuthenticateFunc
	keeper           jwt.SecretsKeeper
	strategy         union.Union
	cache            libcache.Cache
}

// New Guardian factory
func New(authenticateFunc basic.AuthenticateFunc, keeper jwt.SecretsKeeper) *Guardian {

	cache := libcache.FIFO.New(0)
	cache.SetTTL(time.Minute * 5)
	cache.RegisterOnExpired(func(key, _ interface{}) {
		cache.Peek(key)
	})

	basicStrategy := basic.NewCached(authenticateFunc, cache)
	jwtStrategy := jwt.New(cache, keeper)
	strategy := union.New(jwtStrategy, basicStrategy)

	return &Guardian{
		keeper:    keeper,
		strategy:  strategy,
		cache:     cache,
		whitelist: map[string]string{"/users": "POST"},
	}
}

// CreateToken return a JWT token
func (am *Guardian) CreateToken(w http.ResponseWriter, r *http.Request) {
	u := auth.User(r)
	u.GetExtensions().Set("x-go-guardian-basic-password", "jwt")
	token, _ := jwt.IssueAccessToken(u, am.keeper)
	body := fmt.Sprintf(`{"token":"%s"}`, token)
	w.Write([]byte(body))
}

// AuthMiddleware Authentication middleware
func (am *Guardian) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, ok := am.whitelist[r.URL.Path]
		if ok && r.Method == method {
			next.ServeHTTP(w, r)
			return
		}

		log.Println("Executing Auth Middleware")
		_, user, err := am.strategy.AuthenticateRequest(r)
		if err != nil {
			fmt.Println(err)
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		log.Printf("User %s Authenticated\n", user.GetUserName())
		r = auth.RequestWithUser(user, r)
		next.ServeHTTP(w, r)
	})
}

// GetUserID return User ID
func (am *Guardian) GetUserID(r *http.Request) string {
	u := auth.User(r)
	return u.GetID()
}
