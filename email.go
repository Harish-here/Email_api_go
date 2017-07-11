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

func getTemplates(w http.ResponseWriter, r *http.Request){
	t,_ := template.ParseFiles("confirm.html")//parsing the file to template
	buf := new(bytes.Buffer) //contanier to hold the string of html code
	err := t.Execute(buf,person{customerName:"Madhan",phoneNumber:"9791528548",customerEmail:"madhan@infonixweblab.com",message:"hi this is a sample"})
	checkerr(err)
	fmt.Fprintln(w,buf)

}

func sendMail(w http.ResponseWriter,from string,to string,subject string,body string){
	//main function for sending email
	m := gomail.NewMessage()
	m.SetHeader("From",from)
	m.SetHeader("To",to)
	m.SetHeader("subject",subject)
	m.SetBody("text/html",body)
	d := gomail.NewDialer("smtp.zoho.com", 587, "harish@infonixweblab.com", "harish123")
	if err := d.DialAndSend(m); err != nil {
   		 panic(err)
		}
			fmt.Fprint(w,"mail sent")
		
}



func getTemplate(file string,data templatedata) string {
	//function to make body of the email (template + Data)
	t,_ := template.ParseFiles(file)
	buf := new(bytes.Buffer)
	if err = t.Execute(buf,data);err != nil {
		panic(err)
	}
		return buf.String()
	
}
type person struct{
	customerName string
	phoneNumber string
	customerEmail string
	message string
}
 
type templatedata struct{
	CustomerName string
	PhoneNumber string
	CustomerEmail string
	Message string
}
func sendMails(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
var hotelID string
var emailType string
	if r.FormValue("hotelId") != "" || r.FormValue("emailType") != ""{ //checking for the postData
		//getting the hotel id and it's type
			hotelID = r.FormValue("hotelId")
			emailType = r.FormValue("emailType") 
		
	}else{
		hotelID ="h2"
		emailType = "2"
	}
	//mockData
	d:= templatedata{
		CustomerName : "harsih",
		PhoneNumber :"8870072364",
		CustomerEmail : " harish@infonixweblab.com",
		Message : " this is a sample message",
	}
	var patternName string
//get the pattern from db for respective hotel based on hotelId

 	rows ,err := db.Query("SELECT patternName from hotel_email_pattern WHERE hoteId =? AND emailType =?",hotelID,emailType)
 	checkerr(err)
	 
	for rows.Next(){
	
	err:=rows.Scan(&patternName)
	checkerr(err)

	}
	body := getTemplate(patternName,d)
	sendMail(w,"harish@infonixweblab.com","madhan@infonixweblab.com","Booking Confirmation",body)
}
func main(){
	db, err = sql.Open("mysql", "root:@/user_db")
	checkerr(err)
	defer db.Close()



	http.HandleFunc("/people",something)
	http.HandleFunc("/",sayhi)
	http.HandleFunc("/template",getTemplates)
	http.HandleFunc("/sendmail",sendMails)
	log.Fatal(http.ListenAndServe(":8000",nil))
}