package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	apiUrl := "https://randomuser.me/api/"

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Println(err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	w.WriteHeader(http.StatusOK)
	w.Write(data)

}
