package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custome form struct and embeds a url.value field
type Form struct {
	url.Values
	Errors errors
}

// New initializes a for struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	if x == "" {
		return false
	} else {
		return true
	}
}
