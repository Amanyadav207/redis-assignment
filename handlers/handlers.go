package handlers

import (
	"encoding/json"
	"key-value-cache/cache"
	"net/http"
)

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
}

var cacheInstance = cache.NewCache()

func PutHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"status": "ERROR", "message": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	if len(req.Key) > 256 || len(req.Value) > 256 {
		http.Error(w, `{"status": "ERROR", "message": "Key or value exceeds 256 characters"}`, http.StatusBadRequest)
		return
	}

	cacheInstance.Put(req.Key, req.Value)
	json.NewEncoder(w).Encode(Response{Status: "OK", Message: "Key inserted/updated successfully."})
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if len(key) > 256 {
		http.Error(w, `{"status": "ERROR", "message": "Key exceeds 256 characters"}`, http.StatusBadRequest)
		return
	}

	if value, found := cacheInstance.Get(key); found {
		json.NewEncoder(w).Encode(Response{Status: "OK", Key: key, Value: value})
	} else {
		http.Error(w, `{"status": "ERROR", "message": "Key not found."}`, http.StatusNotFound)
	}
}