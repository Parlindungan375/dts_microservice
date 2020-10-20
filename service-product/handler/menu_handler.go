package handler

import (
	"net/http"

	"github.com/Parlindungan375/dts_microservice/utils"
)

// AddMenuHandler handle add menu
func AddMenuHandler(w http.ResponseWriter, r *http.Request) {

	utils.WrapAPISuccess(w, r, "success", 200)
}
