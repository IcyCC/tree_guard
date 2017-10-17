package model

import (
	"database/sql"
	"errors"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"crypto/md5"
	"fmt"
	"math/rand"
)


var(
	DB,_ = sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/tree_guard?charset=utf8")
)


type Tree struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Hash string `json:"hash"`
	Height float32 `json:"height"`
}


func NewTree(name string, hash string, height float32 ) (*Tree, error) {
	return &Tree{
		Name:name,
		Hash:hash,
		Height:height,
	},nil
}

func FindTreeById(id int) (*Tree, error){

	var t Tree
	err := DB.QueryRow("select  * from tree WHERE id = ?",id).Scan(&t.Id, &t.Name, &t.Hash, &t.Height)

	if err != nil{
		return nil,errors.New("no result")
	}

	return &t, nil
}


func FindTreeByHash(hash string) (*Tree, error){

	var t Tree
	err := DB.QueryRow("select  * from tree WHERE hash = ?",hash).Scan(&t.Id, &t.Name, &t.Hash, &t.Height)

	if err != nil{
		return nil,errors.New("no result")
	}

	return &t, nil
}

func GetAllTrees(hash string) ([]*Tree, error){
	data := []*Tree{}
	rows, err := DB.Query("select  * from tree")

	if err != nil{
		log.Println(err)
		return data,err
	}

	for rows.Next(){
		var t Tree
		rows.Scan(&t.Id,&t.Name,&t.Hash,&t.Height)
		data = append(data, &t)
	}
	return data,nil
}



func (tree *Tree) SaveOrUpdate() error{
	res,err := DB.Exec("INSERT INTO tree(name,hash,height) VALUES (?,?,?) ON DUPLICATE KEY UPDATE name = VALUES(name),height=VALUES(height)",tree.Name, tree.Hash, tree.Height)
	if err!=nil{
		log.Println("Tree id : ", tree.Id, " save fail")
		return err
	}
	log.Println("insert success eff id : ", res)
	return nil
}

func (tree *Tree) Update() error{
	res,err := DB.Exec("INSERT INTO tree(id,name,hash,height) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE name = VALUES(name),height=VALUES(height)",tree.Name, tree.Hash, tree.Height)
	if err!=nil{
		log.Println("Tree id : ", tree.Id, " save fail")
		return err
	}
	log.Println("insert success eff id : ", res)
	return nil
}

func GenHash( salt string, args ...string)(string){
	rand.Seed(time.Now().UnixNano())

	var res string

	for _, arg := range args {
		res = res+arg
	}
	res =salt+res
	s := md5.Sum([]byte(res))
	res = fmt.Sprintf("%x",s)
	return res
}

