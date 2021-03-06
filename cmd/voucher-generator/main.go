package main

import (
	"flag"
	"github.com/golang/glog"
	"net/http"

	"github.com/gregbiv/voucher-generator/pkg/controllers/web"
	"github.com/gregbiv/voucher-generator/pkg/system"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	gojiweb "github.com/zenazn/goji/web"
)

func main() {
	filename := flag.String("config", "config.json", "Path to configuration file")

	flag.Parse()
	defer glog.Flush()

	var application = &system.Application{}

	application.Init(filename)
	application.LoadTemplates()

	// Setup static files
	static := gojiweb.New()
	static.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir(application.Configuration.PublicPath))))

	http.Handle("/assets/", static)

	// Apply middleware
	goji.Use(application.ApplyTemplates)
	goji.Use(application.ApplySessions)

	controller := &web.Controller{}

	// Couple of files - in the real world you would use nginx to serve them.
	goji.Get("/robots.txt", http.FileServer(http.Dir(application.Configuration.PublicPath)))
	goji.Get("/favicon.ico", http.FileServer(http.Dir(application.Configuration.PublicPath+"/images")))

	// Home page
	goji.Get("/", application.Route(controller, "Index"))

	graceful.PostHook(func() {
		application.Close()
	})
	goji.Serve()
}
