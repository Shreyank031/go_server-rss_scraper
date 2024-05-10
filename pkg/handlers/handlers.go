package handlers

import (
	"net/http"

	"github.com/Shreyank031/go-rss_scraper/utils"
)

func HandlerReadiness(w http.ResponseWriter, h *http.Request) {
	utils.RespondWithJson(w, http.StatusOK, struct{}{})

}
func HandlerError(w http.ResponseWriter, h *http.Request) {

	utils.RespondWithError(w, 400, "Something went wrong!")

}
