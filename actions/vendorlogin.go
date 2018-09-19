package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func Vendorlogin(c buffalo.Context) error {

	return c.Render(200, r.HTML("Vendor_Login.html", "master_page/login_tmp.html"))
}
