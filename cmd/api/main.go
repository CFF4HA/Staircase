package main

import (
	"flag"
	"net/http"

	"github.com/CFF4HA/Staircase/internal/api/database"
	routes "github.com/CFF4HA/Staircase/internal/api/http/router"

	"github.com/DAlba-sudo/pbf"
)

func main() {
	fAddress := flag.String("address", "", "the address for the scraping http server to bind to")
	fFrontendOrigin := flag.String("frontend", "http://127.0.0.1:8081", "the address for the backend api")
	fPort := flag.Int("port", 8000, "the port for the scraping http server to bind to")
	fDatabaseUrlScraping := flag.String("db-scraping", "", "the database connection url for the scraping database")
	fCertFile := flag.String("cert", "", "the path to the TLS certificate file for the scraping http server")
	fKeyFile := flag.String("key", "", "the path to the TLS key file for the scraping http server")
	flag.Parse()

	database.InitDatabase(*fDatabaseUrlScraping)

	router := pbf.CreateRouter()
	router.Address = *fAddress
	router.Port = *fPort
	router.CertificateFile = *fCertFile
	router.KeyFile = *fKeyFile

	if *fFrontendOrigin != "" {
		routes.SetBackendUrl(*fFrontendOrigin)
	}
	router.SetMiddleware(routes.CORS)

	router.Add(pbf.RouteOptions{
		Endpoint: "/job",
		Method:   http.MethodPut,
		Handler:  routes.HandleJobPUT,
	})

	router.Add(pbf.RouteOptions{
		Endpoint: "/job",
		Method:   http.MethodGet,
		Handler:  routes.HandleJobGET,
	})

	router.Add(pbf.RouteOptions{
		Endpoint: "/job",
		Method:   http.MethodDelete,
		Handler:  routes.HandleJobDELETE,
	})
	router.Add(pbf.RouteOptions{
		Endpoint: "/job",
		Method:   http.MethodPatch,
		Handler:  routes.HandleJobPATCH,
	})

	// User related routes
	router.Add(pbf.RouteOptions{
		Endpoint: "/user",
		Method:   http.MethodPut,
		Handler:  routes.HandleUserPUT,
	})

	router.Add(pbf.RouteOptions{
		Endpoint: "/user",
		Method:   http.MethodPost,
		Handler:  routes.HandleUserPOST,
	})

	// Datasource related routes
	router.Add(pbf.RouteOptions{
		Endpoint: "/ds",
		Method:   http.MethodPut,
		Handler:  routes.HandleDatasourcePUT,
	})

	if err := router.Start(); err != nil {
		panic(err)
	}
}
