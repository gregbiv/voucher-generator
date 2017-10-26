package web

import (
	"net/http"

	"github.com/gregbiv/voucher-generator/pkg/helpers"
	"github.com/zenazn/goji/web"
)

// Home page route
func (controller *Controller) Index(c web.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(c)

	// With that kind of flags template can "figure out" what route is being rendered
	c.Env["IsIndex"] = true

	c.Env["Voucher"] = helpers.GenerateVoucher()

	c.Env["Title"] = "Voucher generator"

	return helpers.Parse(t, "main", c.Env), http.StatusOK
}
