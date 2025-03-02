package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Home struct {
	someThingDate []int
}

type AddNumberData struct {
	Number int `json:"number"`
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		for i := 0; i < len(h.someThingDate); i++ {
			fmt.Fprintf(w, "%d >> %#v \n", i, h.someThingDate[i])
		}

	case http.MethodPost:
		fmt.Fprintf(w, "you chosed post method\n")
		var data AddNumberData
		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			http.Error(w, "invalid request body\n", http.StatusBadRequest)
			return
		}
		h.someThingDate = append(h.someThingDate, data.Number)
		fmt.Fprintf(w, "Succesful added number>> %d\n\n", data.Number)

		for i := 0; i < len(h.someThingDate); i++ {
			fmt.Fprintf(w, "%d >> %#v \n", i, h.someThingDate[i])
		}
	}
}
func main() {

	home := Home{someThingDate: make([]int, 0)}
	http.Handle("/", &home)
	http.ListenAndServe("localhost:8080", nil)
}
