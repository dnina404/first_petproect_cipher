package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint
	Money                 int16
	Avg_grades, happiness float64
	Hobbies               []string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, 200, 0.2, 0.5, []string{"Voleybol", "Dance", "Anime"}}
	// bob.setNewName("Mihail")
	//fmt.Fprintf(w, `<h1>Hello html</h1>
	//<p>dnina404</p>`)
	tmpl, _ := template.ParseFiles("/home/name/work/go/templates/home_page.html")
	tmpl.Execute(w, bob)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}
func test_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test good")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/testpage/", test_page)
	http.ListenAndServe(":8080", nil)
}
func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is %s. He is %d and he has money "+
		"equal: %v", u.Name, u.Age, u.Money)
}
func main() {
	handleRequest()
}
