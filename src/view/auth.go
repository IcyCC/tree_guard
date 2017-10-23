package view

import (
	"github.com/IcyCC/go_login"
	"net/http"
	"fmt"
	"model"
	"encoding/json"
	"log"
	"strconv"
)

var(
	LoginConfig = go_login.NewConfig("hello")
	LoginManager = go_login.NewLoginManager(LoginConfig)
)



type UserResult struct {
	Status string `json:"status"`
	Data []*model.User `json:"data"`
	Reason string `json:"reason"`
}


func UserHandle(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		result := UserResult{Status:"",Data:[]*model.User{},Reason:""}
		name := r.FormValue("name")
		id := r.FormValue("id")
		if name==""&&id==""{
			result.Status = "Success"
			users,_ := model.GetAllUsers()
			result.Data = users
			b,_ := json.Marshal(result)
			log.Println(string(b))
			fmt.Fprint(w,string(b))
			return
		}
		if name!=""&&id==""{
			user,err := model.FindUserByName(name)
			if err!=nil{
				result.Status = "Fail"
				result.Reason = "no this name user"
				b,_ := json.Marshal(result)
				log.Println(string(b))
				fmt.Fprint(w,string(b))
				return
			}
			result.Data = append(result.Data,user)
			b,_ := json.Marshal(result)
			log.Println(string(b))
			fmt.Fprint(w,string(b))
			return
		}
		if name==""&&id!=""{
			id,_ := strconv.Atoi(id)
			user,err := model.FindUserById(id)
			if err!=nil{
				result.Status = "Fail"
				result.Reason = "no this id user"
				b,_ := json.Marshal(result)
				log.Println(string(b))
				fmt.Fprint(w,string(b))
				return
			}
			result.Data = append(result.Data,user)
			b,_ := json.Marshal(result)
			log.Println(string(b))
			fmt.Fprint(w,string(b))
			return
		}
		if name!=""&&id!=""{
			result.Status = "Fail"
			result.Reason = "can`t both name and id"
			b,_ := json.Marshal(result)
			log.Println(string(b))
			fmt.Fprint(w,string(b))
			return
		}
	}
	if r.Method == "POST"{
		var result UserResult
		name := r.FormValue("name")
		password := r.FormValue("password")
		_,err := model.FindUserByName(name)
		if err==nil{
			result.Status = "Fail"
			result.Reason = "this name already regist"
			b,_ := json.Marshal(result)
			log.Println(string(b))
			fmt.Fprint(w,string(b))
			return
		}
		user := &model.User{Name:name, Password:model.GenPassword(password)}
		user.SaveOrUpdate()
		result.Status = "Success"
		result.Reason = ""
		result.Data = append(result.Data,user)
		b,_ := json.Marshal(result)
		log.Println(string(b))
		fmt.Fprint(w,string(b))
		return
	}
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		fmt.Fprint(w,"<html><head><title></title></head><body><form action=\"http://127.0.0.1:8080/login\" method=\"post\">用户名:<input type=\"text\" name=\"username\">密码:<input type=\"password\" name=\"password\"><input type=\"submit\" value=\"登陆\"></form></body></html>")
	}
	if r.Method == "POST" {
		var result UserResult
		name := r.FormValue("username")
		password := r.FormValue("password")

		if name == "" || password == "" {
			result.Status = "Fail"
			result.Reason = "password or name is empty"
			b, _ := json.Marshal(result)
			log.Println(string(b))
			fmt.Fprint(w, string(b))
			return
		}
		user, err := model.FindUserByName(name)
		if err != nil {
			result.Status = "Fail"
			result.Reason = "no this user"
			b, _ := json.Marshal(result)
			log.Println(string(b))
			fmt.Fprint(w, string(b))
			return
		}

		currentUser,ok := LoginManager.Auth(r)
		if currentUser != nil{
			myUser,_ := currentUser.(*model.User)
			if myUser.Name == user.Name&&ok{
				result.Status = "Fail"
				result.Reason = "User already login"
				b, _ := json.Marshal(result)
				log.Println(string(b))
				fmt.Fprint(w, string(b))
				return
			}
		}

		if !model.VerfiyPassword(password, user.Password) {
			result.Status = "Fail"
			result.Reason = "invalid password"
			b, _ := json.Marshal(result)
			log.Println(string(b))
			fmt.Fprint(w, string(b))
			return
		}
		LoginManager.Login(user,&w)
		result.Status = "Success"
		result.Reason = ""
		result.Data = append(result.Data,user)
		b, _ := json.Marshal(result)
		log.Println(string(b))
		fmt.Fprint(w, string(b))
	}

}



func LogoutHandle(w http.ResponseWriter, r *http.Request) {
	user,ok := LoginManager.Current(r)
	if user == nil || ok == false{
		fmt.Fprint(w,"No one logout")
		return
	}
	LoginManager.Logout(user,r, &w)
	fmt.Fprint(w,"logout")
}

func UserTestHandle(w http.ResponseWriter, r *http.Request){
	currentUser,ok := LoginManager.Auth(r)
	var result UserResult
	if currentUser != nil && ok{
		myUser,_ := currentUser.(*model.User)
		result.Status = "Success"
		result.Reason = ""
		result.Data = append(result.Data,myUser)
		b, _ := json.Marshal(result)
		log.Println(string(b))
		fmt.Fprint(w, string(b))
		return
	} else {
		fmt.Fprint(w,"No one login")
	}

}