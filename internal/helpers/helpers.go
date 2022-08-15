package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/xor0x/bookings/internal/config"
)

var app *config.AppConfig

// sets up app config for helpers
func NewHlpers(a *config.AppConfig) {
	app = a
}


func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status 0", status)
	http.Error(w, http.StatusText(status), status)

}


func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}