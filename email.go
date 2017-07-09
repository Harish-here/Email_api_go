package main
import (
	"net/http"
	"fmt"
	"log"
	"html/template"

	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"bytes"

	"gopkg.in/gomail.v2"
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
	t,_ := template.ParseFiles("confirm.html")//parsing the file to template
	buf := new(bytes.Buffer) //contanier to hold the string of html code
	err := t.Execute(buf,person{Username:"harish"})
	checkerr(err)
	fmt.Fprintln(w,buf)

}

func sendMail(w http.ResponseWriter, r *http.Request){
	m := gomail.NewMessage()
m.SetHeader("From", "harish@infonixweblab.com")
m.SetHeader("To", "justinsylas@infonixweblab.com")
m.SetHeader("Subject", "Hello this is a test mail! from go API")
	t,_ := template.ParseFiles("confirm.html")
	buf := new(bytes.Buffer)
	if err = t.Execute(buf,person{Username:"harish"});err != nil {
		panic(err)
	}
 body := buf.String()
m.SetBody("text/html", body)


d := gomail.NewDialer("smtp.zoho.com", 587, "harish@infonixweblab.com", "harish123")

// Send the email to the target thru host
if err := d.DialAndSend(m); err != nil {
    panic(err)
}else{
	fmt.Fprint(w,"mail sent")
}
}
func main(){
	http.HandleFunc("/people",something)
	http.HandleFunc("/",sayhi)
	http.HandleFunc("/template",getTemplate)
	http.HandleFunc("/sendmail",sendMail)
	log.Fatal(http.ListenAndServe(":3306",nil))
}