package main

import (
	"fmt"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		log.Printf("Platform: %s", cfg.platform)
		w.WriteHeader(http.StatusForbidden)
		w.Write(fmt.Append(nil, "Reset is only allowed in dev environment."))
		return
	}

	cfg.fileserverHits.Store(0)
	err := cfg.db.Reset(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "Failed to reset the database: %s\n", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(fmt.Append(nil, "Hits reset to 0 and database reset to initial state."))
}
