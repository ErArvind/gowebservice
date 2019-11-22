package main 

import (
	"fmt"
	"net/http"
	"github.com/ErArvind/gowebservice/controllers"
	//"github.com/ErArvind/gowebservice/models"
)
func main(){
fmt.Println("main is called")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000",nil)
}