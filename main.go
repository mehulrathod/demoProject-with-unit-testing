package main

//type Author struct {
//	Name string `json:"name"`
//	Age int `json:"age"`
//}
//
//func main() {
//
//	json, err := json.Marshal(Author{Name: "Vijay", Age: 25})
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	err = redis2.SetValue("abc",json)
//	if err != nil {
//		fmt.Println("error", err)
//	}
//	res :=redis2.GetValue("abc")
//	fmt.Println(res)
//}

//package main
//
import (
	"log"
	"os"
	Routers "vipulsirdemo/routers"

	"github.com/joho/godotenv"
)

//Execution starts from main function
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("ERROR : ", err)
	}
	r := Routers.SetupRouter()

	port := os.Getenv("port")

	// For run on requested port
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080" //localhost
	}
	type Job interface {
		Run()
	}

	err = r.Run(":" + port)
	if err != nil {
		log.Println("ERROR")
	}



}
