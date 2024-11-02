package pkg

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status,omitempty"`
	Data   any    `json:"data,omitempty"`
	Error  any    `json:"error,omitempty"`
}

func SendErrResp(w http.ResponseWriter, err error) {
	if customErr, ok := err.(CustomError); ok {
		byteRes, err := json.Marshal(&Response{
			Status: "error",
			Error:  customErr.Message,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(customErr.StatusCode)
		w.Write(byteRes)
		return
	}

	byteRes, err := json.Marshal(&Response{
		Status: "error",
		Error:  ErrInternal.Message,
	})
	if err != nil {
		w.WriteHeader(ErrInternal.StatusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ErrInternal.StatusCode)
	w.Write(byteRes)
}

func SendSuccessOkResp(w http.ResponseWriter, resp any) {
	byteRes, err := json.Marshal(&Response{
		Status: "success",
		Data:   resp,
	})
	if err != nil {
		w.WriteHeader(ErrInternal.StatusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteRes)
}

func SendSuccessCreatedResp(w http.ResponseWriter, resp any) {
	byteRes, err := json.Marshal(&Response{
		Status: "success",
		Data:   resp,
	})
	if err != nil {
		w.WriteHeader(ErrInternal.StatusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(byteRes)
}
