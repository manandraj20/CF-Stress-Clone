package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

func home(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"Test case#1 passed!",
			"Test case#2 passed!",
		},
	}

	err = t.Execute(w, data)
	check(err)

	path := r.Method

	fmt.Fprintf(w, "You requested for the path: %s\n", path)
	fmt.Fprintf(w, "Hello!..You are in Home Page")
}
func pricing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
	fmt.Fprintf(w, "Pricing!")

}
func status(w http.ResponseWriter, r *http.Request) {
	var outputStreamBuffer, errorStreamBuffer bytes.Buffer
	cmd := exec.Command("/usr/bin/bash", "/home/manandraj20/dummyScript.sh")
	cmd.Stdout = &outputStreamBuffer
	cmd.Stderr = &errorStreamBuffer
	err := cmd.Run()
	fmt.Fprintf(w, "output buffer was :%s\n", outputStreamBuffer.String())
	fmt.Fprintf(w, "error buffer was: %s\n", errorStreamBuffer.String())
	if err != nil {
		log.Fatal(err)
	}
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Ticket: %v\n", vars["ticket"])
	fmt.Fprintf(w, "Profile!")

}
func mailingList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mailing list!")
}
func blog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Blog!")
}
func getMoreSamples(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "ContestID: %v\n", vars["contestID"])
	fmt.Fprintf(w, "Problem Index: %v\n", vars["problemIndex"])
}

func main() {
	r := mux.NewRouter()
	s := r.Host("localhost:4000").Subrouter()
	//Matches a dynamic subdomain
	//r.Host("{subdomain:[a-z]+}.cfstress.com")

	s.HandleFunc("/", home)
	s.HandleFunc("/test/{contestID:[0-9]+}/{problemIndex:[a-z]+}/", getMoreSamples)
	s.HandleFunc("/Pricing/{category}", pricing)

	s.HandleFunc("/MailingList", mailingList)
	s.HandleFunc("/Blog", blog)
	s.HandleFunc("/Status/{ticket:[0-9]+}", status)
	http.Handle("/", r)
	//http.HandleFunc("/", home)
	//http.HandleFunc("/Pricing", pricing)
	//http.HandleFunc("/Status", status)
	//http.HandleFunc("/MailingList", mailingList)
	//http.HandleFunc("/Blog", blog)
	fmt.Printf("Starting server at port 4000\n")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}

	//log.Println("Printing on the stderr file descriptor")
	//fmt.Println("I think file couldn't save issue has been resolved!")
	//log.Fatal(errors.New("printing to stderr file"))
}
