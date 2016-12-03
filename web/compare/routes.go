package compare

import (
	"github.com/pressly/chi"
)

func Routes(r chi.Router) {
	r.Get("/", compareHandler)
}
