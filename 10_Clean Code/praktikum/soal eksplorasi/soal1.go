package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Request adalah struktur data untuk merepresentasikan request
type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

// Response adalah struktur data untuk merepresentasikan response
type Response struct {
	Result int `json:"result"`
}

// CalculatorController adalah kontroler untuk operasi perhitungan
type CalculatorController struct{}

// AddHandler adalah handler untuk operasi penjumlahan
func (cc *CalculatorController) AddHandler(c echo.Context) error {
	rq := new(Request)
	if err := c.Bind(rq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	return c.JSON(http.StatusOK, Response{Result: rq.A + rq.B})
}

// SubtractHandler adalah handler untuk operasi pengurangan
func (cc *CalculatorController) SubtractHandler(c echo.Context) error {
	rq := new(Request)
	if err := c.Bind(rq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	return c.JSON(http.StatusOK, Response{Result: rq.A - rq.B})
}

// MultiplyHandler adalah handler untuk operasi perkalian
func (cc *CalculatorController) MultiplyHandler(c echo.Context) error {
	rq := new(Request)
	if err := c.Bind(rq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	return c.JSON(http.StatusOK, Response{Result: rq.A * rq.B})
}

// DivideHandler adalah handler untuk operasi pembagian
func (cc *CalculatorController) DivideHandler(c echo.Context) error {
	rq := new(Request)
	if err := c.Bind(rq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if rq.B == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "division by zero"})
	}

	return c.JSON(http.StatusOK, Response{Result: rq.A / rq.B})
}

func main() {
	e := echo.New()

	cc := &CalculatorController{}

	e.POST("/api/add", cc.AddHandler)
	e.POST("/api/subtract", cc.SubtractHandler)
	e.POST("/api/multiply", cc.MultiplyHandler)
	e.POST("/api/divide", cc.DivideHandler)

	e.Start(":1323")
}
