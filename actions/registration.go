package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func Registration(c buffalo.Context) error {

	return c.Render(200, r.HTML("registration_form.html", "master_page/registration_tmp.html"))
}
