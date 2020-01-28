package handler

import (
	"encoding/json"
	"fmt"
	"github.com/tjandrayana/lua-project-v1/go/services"
	"io/ioutil"
	"net/http"
	"log"

)


type blacklistCheckHandler struct {
	service *services.BlacklistModule
}

func NewBlacklistHandler(s *services.BlacklistModule) *blacklistCheckHandler{
	return &blacklistCheckHandler{
		service: s,
	}
}

func (m blacklistCheckHandler)blacklistCheck(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Something Wrong !!!")
		return
	}

	var bp services.BlacklistParameter
	if err := json.Unmarshal(body,&bp); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Something Wrong !!!")
		return
	}

	isBlacklist := m.service.CheckBlacklist(bp)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf(`{"is_blacklist":%t}`,isBlacklist))
	return

}
