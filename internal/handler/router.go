package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Dimoonevs/calculate/factorial/internal/models"
	"github.com/Dimoonevs/calculate/factorial/internal/service"
	"github.com/julienschmidt/httprouter"
)

type AppFactorial struct {
	Service service.ServiceInterface
}

func NewAppFactorial(service service.ServiceInterface) *AppFactorial {
	return &AppFactorial{
		Service: service,
	}
}

func (a *AppFactorial) NewRouter() http.Handler {
	router := httprouter.New()

	router.POST("/calculate", a.meddleware(a.calculate))

	return router
}

func (a *AppFactorial) meddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data := &models.JsonPayload{}
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(data)

		if err != nil || data.A == nil || data.B == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			out, err := json.Marshal(map[string]string{"error": "incorrect input"})
			if err != nil {
				log.Printf("Server stopped with error: %v", err)
				os.Exit(1)
				return
			}
			w.Write(out)
			return
		}
		ctx := context.WithValue(r.Context(), "data", data)
		next(w, r.WithContext(ctx), ps)
	}
}

func (a *AppFactorial) calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data, ok := r.Context().Value("data").(*models.JsonPayload)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		out, err := json.Marshal(map[string]string{"error": "incorrect input"})
		if err != nil {
			log.Printf("Server stopped with error: %v", err)
			os.Exit(1)
			return
		}
		w.Write(out)
		return
	}
	factorialA, factorialB := a.Service.CalculateFactorial(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	out, err := json.Marshal(map[string]string{fmt.Sprintf("%d!", *data.A): factorialA.String(), fmt.Sprintf("%d!", *data.B): factorialB.String()})
	if err != nil {
		log.Printf("Server stopped with error: %v", err)
		os.Exit(1)
		return
	}

	w.Write(out)
}
