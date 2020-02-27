package receber

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func createDB() (db *sqlx.DB, err error) {
	return sqlx.Connect("postgres", "user=horse dbname=white password=whitehorse sslmode=disable")
}

//HandlerReceberCreate ---
func HandlerReceberCreate(c echo.Context) error {
	data := new(Receber)
	err := c.Bind(data)
	if err != nil {
		return err
	}
	result, err := receberCreate(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}

//HandlerReceberRead ---
func HandlerReceberRead(c echo.Context) error {
	result, err := receberRead()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}

//HandlerReceberUpdate ---
func HandlerReceberUpdate(c echo.Context) error {
	data := new(Receber)
	err := c.Bind(data)
	if err != nil {
		return err
	}
	result, err := receberUpdate(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}

//HandlerReceberRemove ---
func HandlerReceberRemove(c echo.Context) error {
	data := new(Receber)
	err := c.Bind(data)
	if err != nil {
		return err
	}
	result, err := receberRemove(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result)
}
