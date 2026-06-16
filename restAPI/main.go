package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age int
	Address string
}

type errorResp struct {
	Statuscode int
	Message string
}

var (
	users = map[string]User{}                                                                    //global variable to store user data in memory
)

func main(){

	http.HandleFunc("/createusers", adduser)
	fmt.Println("user created successfully : ", users)

	fmt.Println("server is running on port 8000")

	log.Fatal("server error , err : %v\n",http.ListenAndServe(":8000", nil))


}
//create/Add  new record of user in map
func adduser(w http.ResponseWriter , r *http.Request){                                         //check if the request method is POST, if not return 405 Method Not Allowed
   if r.Method != "POST"{
	  w.WriteHeader(http.StatusMethodNotAllowed)

	  //create error response struct with status code and message 
	  err := errorResp{                                                                      //create error response struct with status code and message
			Statuscode: http.StatusBadRequest,
			Message: "Method not allowed",
		}
		json.NewEncoder(w).Encode(err)
	  return
    }
	user := User{}                                                                            //create empty user struct to hold the data from request body
	err :=json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		err := errorResp{                                                                      //create error response struct with status code and message
			Statuscode: http.StatusBadRequest,
			Message: "Invalid request body"+err.Error(),
		}
		json.NewEncoder(w).Encode(err)
		return
	}
	users[user.Name] = user                                                                             //store user data in map with name as key
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
	
	fmt.Println("user created successfully :", users)
 return
}
