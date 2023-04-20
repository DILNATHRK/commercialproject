package portal

import (
	// "html/template"

	// "net/http"

	"commercial-propfloor/controller/commercial/manage_commercial"
	"fmt"

	"github.com/gin-gonic/gin"
)

// type Prop struct {
// 	Title     string
// 	MainImage string
// 	Done      bool
// }

// type PropFloor struct {
// 	PageTitle string
// 	Props     []Prop
// }

// func Ankit() string   {
// 	a := "ankit"
// 	return a
// }

func Ankit() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := "ankit"
		fmt.Println(a)
		return
	}
}

// func Cities() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Sprint("HI",manage_commercial.SelectCityindatabase())
// 		return
// 	}
// }
