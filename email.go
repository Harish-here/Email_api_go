package main
import (

	"net/http"
	"log"
	"fmt"
)
func something(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"hello")
}

func main(){
	http.HandleFunc("/",something)
	log.Fatal(http.ListenAndServe(":8000",nil))
}