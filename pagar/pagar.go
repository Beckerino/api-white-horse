package pagar

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func createDB() (db *sqlx.DB, err error) {
	return sqlx.Connect("postgres", "user=horse dbname=white password=whitehorse sslmode=disable")
}

//HandlerPagarCreate ---
func HandlerPagarCreate(c echo.Context) error {
	data := new(Receber)
	err := c.Bind(data)
	if err != nil {
		return err
	}
	result, err := pagarCreate(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}

//HandlerPagarRead ---
func HandlerPagarRead(c echo.Context) error {
	result, err := pagarRead()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}

//HandlerPagarUpdate ---
func HandlerPagarUpdate(c echo.Context) error {
	data := new(Receber)
	err := c.Bind(data)
	if err != nil {
		return err
	}
	result, err := pagarUpdate(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}

//HandlerPagarRemove ---
func HandlerPagarRemove(c echo.Context) error {
	data := new(Receber)
	err := c.Bind(data)
	if err != nil {
		return err
	}
	result, err := pagarRemove(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}
