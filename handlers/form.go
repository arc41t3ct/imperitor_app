package handlers

import (
	"fmt"
	"imperatorapp/models"
	"net/http"

	jet "github.com/CloudyKit/jet/v6"
)

// Form handles request for our form validation form test
func (h *Handlers) Form(w http.ResponseWriter, r *http.Request) {
	variables := make(jet.VarMap)
	validator := h.App.Validator
	variables.Set("validator", validator)
	variables.Set("user", models.User{})
	if err := h.App.Render.Page(w, r, "form", variables, nil); err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) FormPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.App.ErrorLog.Println(err)
		return
	}

	validator := h.App.GetValidator()
	validator.Required(r, "first_name", "last_name", "email")
	validator.IsEmail("email", r.Form.Get("email"))
	validator.Check(len(r.Form.Get("first_name")) > 1, "first_name", "Must be at least two characters long")
	validator.Check(len(r.Form.Get("last_name")) > 1, "last_name", "Must be at least two characters long")

	if !validator.Valid() {
		vars := make(jet.VarMap)
		vars.Set("validator", validator)
		var user models.User
		user.FirstName = r.Form.Get("first_name")
		user.LastName = r.Form.Get("last_name")
		user.Email = r.Form.Get("email")
		vars.Set("user", user)
		if err := h.App.Render.Page(w, r, "form", vars, nil); err != nil {
			h.App.ErrorLog.Println(err)
			return
		}
		return
	}

	fmt.Fprint(w, "valid data")
}
