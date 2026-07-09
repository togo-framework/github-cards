// Package githubcards is the GitHub Cards togo plugin: Cached GitHub proxy: normalized repo stars/forks/language/topics/license/og-image and README with relative links absolutized, stale-on-error fallback. Powers open-source project cards.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package githubcards

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("github-cards", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/github-cards/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"github-cards","status":"ok"}`))
		})
		return nil
	})
}
