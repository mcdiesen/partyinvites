package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// responses new slice of Rsvp struct
var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	//TODO - load templates here
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}

	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

// welcomeHandler function
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// 	...
	templates["welcome"].Execute(w, nil)
}

// listHandler function
func listHandler(w http.ResponseWriter, r *http.Request) {
	// 	...
	templates["list"].Execute(w, responses)
}

type formData struct {
	*Rsvp
	Errors []string
}

// formHandler function
func formHandler(w http.ResponseWriter, r *http.Request) {
	// 	...Execute the template form.html
	if r.Method == http.MethodGet {
		templates["form"].Execute(w, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		responseData := Rsvp{
			Name:       r.Form["Name"][0],
			Email:      r.Form["Email"][0],
			Phone:      r.Form["Phone"][0],
			WillAttend: r.Form["WillAttend"][0] == "true",
		}
		//Validate the fields of the form
		errors := []string{}
		if responseData.Name == "" {
			errors = append(errors, "Please enter your name")
		}
		if responseData.Email == "" {
			errors = append(errors, "Please enter your email")
		}
		if responseData.Phone == "" {
			errors = append(errors, "Please enter your phone number")
		}
		if len(errors) > 0 {
			templates["form"].Execute(w, formData{
				Rsvp: &responseData, Errors: errors,
			})
		} else {
			responses = append(responses, &responseData)
			if responseData.WillAttend {
				templates["thanks"].Execute(w, responseData.Name)
			} else {
				templates["sorry"].Execute(w, responseData.Name)
			}
		}
	}
}

// Main function
func main() {
	// 	...
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
