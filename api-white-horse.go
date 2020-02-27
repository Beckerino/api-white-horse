package main

import (
	"net/http"

	"github.com/beckerino/api-white-horse/pagar"
	"github.com/beckerino/api-white-horse/receber"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	p := e.Group("/pagar")
	p.POST("", pagar.HandlerPagarCreate)
	p.GET("", pagar.HandlerPagarRead)
	p.PUT("", pagar.HandlerPagarUpdate)
	p.DELETE("", pagar.HandlerPagarRemove)

	r := e.Group("/receber")
	r.POST("", receber.HandlerReceberCreate)
	r.GET("", receber.HandlerReceberRead)
	r.PUT("", receber.HandlerReceberUpdate)
	r.DELETE("", receber.HandlerReceberRemove)

	e.Logger.Fatal(e.Start(":9000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
