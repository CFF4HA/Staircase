package main

import (
	"flag"

	"github.com/CFF4HA/Staircase/internal/web/bridges"
	"github.com/CFF4HA/Staircase/internal/web/core"
	"github.com/CFF4HA/Staircase/internal/web/http/middleware"
	"github.com/DAlba-sudo/pff"
)

func main() {
	address := flag.String("address", "0.0.0.0", "the address to bind to for the server")
	backend := flag.String("backend", "", "the backend url to use for the server")
	port := flag.Int("port", 8080, "the port to bind to for the server")
	live := flag.Bool("live", false, "whether to enable live reloading of templates and static files")
	template := flag.String("template", "templates", "the directory path for HTML templates")
	static := flag.String("static", "static", "the directory path for static files")
	cert := flag.String("cert", "", "the path to the TLS certificate file for the server")
	key := flag.String("key", "", "the path to the TLS key file for the server")
	flag.Parse()

	core.BackendURL = *backend
	core.TemplateDir = *template

	app := pff.CreateApp(pff.Configuration{
		Address: *address,
		Port:    *port,
		Live:    *live,

		TemplateDirectoryPath: *template,
		FileSystemPath:        *static,
	})

	app.Router().CertificateFile = *cert
	app.Router().KeyFile = *key

	app.AddMiddleware(middleware.MwNoSession)

	// Concrete Pages
	app.RegisterTemplate("", "index/page.html", pff.TemplateRegistrationOpts{
		IncludeBaseTemplate: true,
	})

	app.RegisterTemplate("/login", "user/login/page.html", pff.TemplateRegistrationOpts{
		IncludeBaseTemplate: true,
	})

	// HTMX Forms
	form_login := app.RegisterTemplate("/htmx/form/login", "user/form/login.html", pff.TemplateRegistrationOpts{
		IncludeBaseTemplate: false,
	})
	_ = form_login

	app.RegisterTemplate("/htmx/form/signup", "user/form/signup.html", pff.TemplateRegistrationOpts{
		IncludeBaseTemplate: false,
	})

	app.RegisterTemplate("/htmx/form/success", "component/form/success.html", pff.TemplateRegistrationOpts{
		IncludeBaseTemplate: false,
	})

	app.RegisterTemplate("/htmx/form/datasource", "datasource/form/create.html", pff.TemplateRegistrationOpts{
		IncludeBaseTemplate: false,
	})

	form_job := app.RegisterTemplate("/htmx/form/job", "job/form/create.html", pff.TemplateRegistrationOpts{
		IncludeBaseTemplate: false,
	})
	form_job.RegisterBridge("Sources", bridges.BridgeDatasources{})

	app.RegisterTemplate("/htmx/form/job/staircase-input", "job/form/staircase-input.html", pff.TemplateRegistrationOpts{})

	// HTMX Datasource
	ds_list := app.RegisterTemplate("/htmx/ds/list", "datasource/card-list.html", pff.TemplateRegistrationOpts{})
	ds_list.RegisterBridge("Sources", bridges.BridgeDatasources{})

	ds := app.RegisterTemplate("/htmx/ds", "datasource/ds.html", pff.TemplateRegistrationOpts{})
	ds.RegisterBridge("Source", bridges.BridgeDatasource{})

	// HTMX Jobs
	job_list := app.RegisterTemplate("/htmx/job/list", "job/card-list.html", pff.TemplateRegistrationOpts{})
	job_list.RegisterBridge("Jobs", bridges.BridgeJobs{})

	job := app.RegisterTemplate("/htmx/job", "job/job.html", pff.TemplateRegistrationOpts{})
	job.RegisterBridge("Job", bridges.BridgeJob{})

	if err := app.Start(); err != nil {
		panic(err)
	}
}
