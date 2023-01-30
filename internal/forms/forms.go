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

// Valid returns true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a for struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if from field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}
