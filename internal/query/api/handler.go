package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (srv *Server) PostHello(c echo.Context) error {
	post := struct {
		Msg *string `json:"msg"`
	}{}

	if err := c.Bind(&post); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if post.Msg == nil {
		return c.String(http.StatusBadRequest, "No data")
	}
	if err := srv.uc.AddHello(*post.Msg); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "")
}

func (srv *Server) GetHellos(c echo.Context) error {
	post := struct {
		Hellos []string `json:"AllValues"`
	}{}
	k, err := srv.uc.AllHellos()
	post.Hellos = k
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}
func (srv *Server) DeleteHello(c echo.Context) error {
	post := struct {
		Msg *string `json:"msg"`
		All *bool   `json:"all"`
	}{}

	if err := c.Bind(&post); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if post.Msg == nil && post.All == nil {
		return c.String(http.StatusBadRequest, "No data")
	}
	if err := srv.uc.DeleteHello(*post.Msg, *post.All); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "")
}
