package main
import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	// "encoding/json"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"os"
)
type userData struct {
    id int
	firstName string
	secondName string
	gender string
}
var db *sql.DB
var err error
func setJSONAsHeader(w http.ResponseWriter){
	w.Header().Set("Content-Type","application/json")
}
func something(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil || len(r.Form) == 0 {
		fmt.Fprintln(w,"No post data")
		return
	}
	// fmt.Fprintln(w,r.Form)
	setJSONAsHeader(w)
	json.NewEncoder(w).Encode(r.Form)
}
func sayhi(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var ser string
	if r.FormValue("id") != ""{
		 ser = r.FormValue("id")
	}
	
	var (
		id string
		firstName string
		secondName string
		gender string
	)
db, err = sql.Open("mysql", "root:@/user_db")
	checkerr(err)
	defer db.Close()
	rows,err := db.Query("SELECT * FROM user_tb WHERE id ="+ser)
	checkerr(err)
	for rows.Next(){
		err := rows.Scan(&id,&firstName,&secondName,&gender) 
		checkerr(err)
		fmt.Fprintln(w,id + " " + firstName +" "+secondName+" "+gender)
	}
	
}
func checkerr(err error){
	if err != nil {
			log.Fatal(err)
		}
}
type person struct{
	Username string
}
func getTemplate(w http.ResponseWriter, r *http.Request){
	t := template.New("sample.html")
	e,_ := t.Parse("hello {{.Username}}")
	p := person{Username:"harish"}
	e.Execute(os.Stdout,p)

}
func main(){
	http.HandleFunc("/people",something)
	http.HandleFunc("/",sayhi)
	http.HandleFunc("/template",getTemplate)
	http.ListenAndServe(":8080",nil)
}