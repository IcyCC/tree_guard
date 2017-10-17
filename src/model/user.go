package model

import (
	"github.com/IcyCC/go_login"
	"log"
	"errors"
)

type User struct {
	go_login.UserMixin
	Id string `json:"id"`
	Name string	`json:"name"`
	Password string
}

func FindUserById(id int) (*User, error){

	var u User
	err := DB.QueryRow("select id,name,password from user WHERE id = ?",id).Scan(&u.Id, &u.Name)

	if err != nil{
		return nil,errors.New("no result")
	}

	return &u, nil
}


func FindUserByName(name string) (*User, error){

	var u User
	err := DB.QueryRow("select  id,name,password from user WHERE name = ?",name).Scan(&u.Id, &u.Name, &u.Password)

	if err != nil{
		return nil,errors.New("no result")
	}

	return &u, nil
}

func GetAllUsers() ([]*User, error){
	data := []*User{}
	rows, err := DB.Query("select id,name,password from user")

	if err != nil{
		log.Println(err)
		return data,err
	}

	for rows.Next(){
		var u User
		rows.Scan(&u.Id,&u.Name,&u.Password)
		data = append(data, &u)
	}
	return data,nil
}



func (user *User) SaveOrUpdate() error{
	res,err := DB.Exec("INSERT INTO user(name,password) VALUES (?,?) ON DUPLICATE KEY UPDATE name = VALUES(name),password=VALUES(password)",user.Name, user.Password)
	if err!=nil{
		log.Println("Tree id : ", user.Id, " save fail")
		return err
	}
	log.Println("insert success eff id : ", res)
	return nil
}

func (user *User) Update() error{
	res,err := DB.Exec("INSERT INTO user(id,name,password) VALUES (?,?,?) ON DUPLICATE KEY UPDATE name = VALUES(name),password=VALUES(password)",user.Name,user.Password)
	if err!=nil{
		log.Println("Tree id : ", user.Id, " save fail")
		return err
	}
	log.Println("insert success eff id : ", res)
	return nil
}

func GenPassword(pw string)string{
	return GenHash("1001", pw)
}

func VerfiyPassword(pw, pw_hash string)bool{
	pw = GenPassword(pw)
	if pw == pw_hash{
		return true
	}
	return false
}