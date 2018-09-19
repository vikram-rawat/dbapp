package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func Vendorpage(c buffalo.Context) error {

	return c.Render(200, r.HTML("Vendor_Page.html", "master_page/Vendor_tmp.html"))
}
