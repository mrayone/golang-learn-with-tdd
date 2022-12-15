package contexts

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return //todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}

// more about context value : https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39
// https://faiface.github.io/post/context-should-go-away-go2/
// https://go.dev/blog/context
