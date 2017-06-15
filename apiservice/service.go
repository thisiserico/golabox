package apiservice

import (
	"log"
	"net/http"

	"github.com/thisiserico/golabox/domain/service"
	"github.com/thisiserico/golabox/readdomain"

	goji "goji.io"
)

type Client struct {
	mux           *goji.Mux
	domainService *service.Service
	readClient    readdomain.Repository
}

func New(srv *service.Service, rc readdomain.Repository) *Client {
	return &Client{
		mux:           goji.NewMux(),
		domainService: srv,
		readClient:    rc,
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (cl *Client) Run() {
	cl.mux.Use(cl.recoverPanic)
	cl.mux.Use(cl.handleNotFoundErrors)
	cl.mux.Use(cl.respondWithJSON)
	cl.mux.Use(cl.logRequest)

	cl.defineMonitoringRoutes()
	cl.defineProductRoutes()
	cl.defineOrderRoutes()
	cl.definePaymentRoutes()
	cl.defineReadModelRoutes()

	log.Fatal(http.ListenAndServe("localhost:8000", cl.mux))
}
