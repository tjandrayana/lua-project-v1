package handler

import (
	"encoding/json"
	"github.com/tjandrayana/lua-project-v1/go/services"
	"io/ioutil"
	"log"
	"net/http"
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
		return
	}

	var bp services.BlacklistParameter
	if err := json.Unmarshal(body,&bp); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isBlacklist := m.service.CheckBlacklist(bp)
	if isBlacklist{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
