package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Space struct {
	Building string `json:"building"`
	Floor string `json:"floor"`
	RoomNr string `json:"roomnr"`
	Name string `json:"name"`
}

var spaces []Space

func main() {

	spaces = []Space {
		{
			Building: "Exactum",
			Floor: "K",
			Name: "Gurula",
		},
		{
			Building: "Physicum",
			Floor: "1",
			RoomNr: "GWR123",
			Name: "Group Study Room (3 persons)",
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/spaces", ListAllSpaces)
	router.HandleFunc("/space/{spaceId}", GetSpace)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Server Works!"))
}

func ListAllSpaces(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "list study spaces")

	b, err := json.Marshal(spaces)
	if err != nil {
		fmt.Println("error when marshaling study spaces")
		fmt.Println(err)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Println("error writing to response")
		fmt.Println(err)
		return
	}
	w.WriteHeader(200)
}

func GetSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studySpaceIdRaw, err := strconv.ParseInt(vars["spaceId"], 10, 32)
	studySpaceId := int(studySpaceIdRaw)
	if err != nil {
		fmt.Println("study space id must be a number")
		w.WriteHeader(404)
		return
	}

	fmt.Println(w, "get space with id: %d", studySpaceId)

	if studySpaceId < 0 || studySpaceId >= len(spaces) {
		fmt.Println("invalid id")
		w.WriteHeader(404)
		return
	}

	space := spaces[studySpaceId]

	b, err := json.Marshal(space)
	if err != nil {
		fmt.Println("error when marshaling study space")
		fmt.Println(err)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Println("error writing to response")
		fmt.Println(err)
		return
	}
	w.WriteHeader(200)
}