package handlers

import (
	"net/http"
	"rest/cmd/api/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		return err
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)

}
func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "unable to process")
	}
	data, err := service.GetById(index)
	if err != nil {
		c.String(http.StatusBadGateway, "unable to process")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)

}
func HandlePost(c echo.Context) error {

	err := service.Postuser(c)
	if err != nil {
		return c.String(http.StatusOK, "unable to process")

	}
	return c.String(http.StatusOK, "user saved successfully")

}
func Removeuser(c echo.Context) error {
	idx := c.Param("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	er := service.DeleteUser(id)
	if er != nil {
		return c.String(http.StatusBadRequest, "user either not exist or cant be deleted")
	}
	return c.String(http.StatusOK, "user deleted")
}
