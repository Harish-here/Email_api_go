package main
import (
	"net/http"
	"fmt"
	// "log"
	// "encoding/json"
)
type userData struct {
    id int
	firstName string
	secondName string
	gender string
}

var err error
func something(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil || len(r.Form) == 0 {
		fmt.Fprintln(w,"No post data")
		return
	}
	fmt.Fprintln(w,r.FormValue("harish"))
}
func sayhi(w http.ResponseWriter, r *http.Request){
fmt.Fprintln(w,"hello")
}

func main(){
	http.HandleFunc("/people",something)
	http.HandleFunc("/",sayhi)
	http.ListenAndServe(":8080",nil)
}