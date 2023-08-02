package auth

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"

	"com.fullstack.ecommerce/app/users/model"
	"com.fullstack.ecommerce/app/users/repository"
	"com.fullstack.ecommerce/utils/helpers"
)

var userCtxKey = &contextKey{"current_user"}

type contextKey struct {
	name string
}

func AuthenticationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			accessToken := header
			username, err := helpers.ValidateToken(accessToken)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			currentUser, err := repository.FindByUsername(r.Context(), username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, currentUser)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}

func Forbidden(ctx context.Context) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: "Forbidden",
		Extensions: map[string]interface{}{
			"code": http.StatusForbidden,
		},
	}
}
