package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

var word Wordslice //global dictionary

//обработка get запроса
func search(w http.ResponseWriter, r *http.Request) {

	fmt.Println(word.Words) // log in consoled
	if r.Method == "GET" {
		anagramResult := Anagramm(r.FormValue("word"), word.Words)
		if anagramResult[0] != "null" {
			fmt.Fprintf(w, "["+strings.Join(anagramResult, ",")+"]") //print to responce
		} else {
			fmt.Fprintf(w, anagramResult[0])
		}
		fmt.Println(anagramResult) // log in console
	}
}

//обработка post запроса
func loader(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &word.Words)

		fmt.Print(word) // loging arr
		if err != nil {
			fmt.Println(err)
		}
	}
}

//SortString is a sort letters in a word
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//Words are a dictionary of words

type Wordslice struct {
	Words []string
}

//Anagramm search anagrams in dictionary Library
func Anagramm(x string, Library []string) []string {
	var Res []string
	for _, s := range Library {
		if len(x) == len(s) {
			if SortString(strings.ToLower(x)) == SortString(strings.ToLower(s)) {
				Res = append(Res, s)
			}
		}
	}
	// проверка на пустой результат
	if Res == nil {
		Res = append(Res, "null")
	}
	return Res

}

func main() {

	http.HandleFunc("/get", search)              // for search (GET) world in list
	http.HandleFunc("/load", loader)             // for load (POST) world list
	log.Fatal(http.ListenAndServe(":8080", nil)) // error log handler
}
