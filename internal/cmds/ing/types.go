package ing

import "net/http"

type (
	gGetter struct {
		c      *http.Client
		rawUrl string
	}

	ingErr struct {
		base    error
		isFatal bool
	}
)

func (h ingErr) Error() string  { return h.base.Error() }
func (h *ingErr) Unwrap() error { return h.base }
func (h *ingErr) Is(target error) bool {
	_, ok := target.(ingErr)
	return ok
}
