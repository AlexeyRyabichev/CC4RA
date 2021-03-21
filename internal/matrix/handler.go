package matrix

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AlexeyRyabichev/CompCompexity4RussiansAlgo/pkg/logger"
)

func CountMatrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var err error
	var rawBody []byte
	if rawBody, err = ioutil.ReadAll(r.Body); err != nil {
		logger.Sugar.Errorf("Error while reading request body. Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var matrices Matrices
	if err = json.Unmarshal(rawBody, &matrices); err != nil {
		logger.Sugar.Errorf("Error while unmarshaling request json. Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := matrices.Validate(); err != nil {
		log.Fatal(err)
	}

	if err := matrices.Prepare(); err != nil {
		logger.Sugar.Errorf("cannot prepare matrices to count. err: %v", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	ans, err := matrices.Multiply()
	if err != nil {
		logger.Sugar.Errorf("cannot miltiply matrices. err: %v", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	var body []byte

	if body, err = json.Marshal(ans); err != nil {
		logger.Sugar.Errorf("Error during json marshaling. Error: %s.", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	if _, err := w.Write(body); err != nil {
		logger.Sugar.Errorf("Error during body write. Error: %s.", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
