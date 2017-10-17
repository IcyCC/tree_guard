package view


import (
	"net/http"
	"fmt"
	"encoding/json"
	"model"
	"log"
	"strconv"
)


type HashResult struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
	Hash string `json:"hash"`
}

type TreeResult struct {
	Status string `json:"status"`
	Data []*model.Tree `json:"data"`
	Reason string `json:"reason"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<h1>Hello Tree</h1>")
}

func TreeHandler(w http.ResponseWriter, r *http.Request){
	resp := TreeResult{Status:"Fail", Data:[]*model.Tree{}}
	if r.Method == "GET"{
		var tree *model.Tree
		var err error
		id := r.FormValue("id")
		hash := r.FormValue("hash")

		if id != ""{
			id, err := strconv.Atoi(id)
			if err != nil{
				resp.Reason = "illegal id"
				b,_ := json.Marshal(resp)
				log.Println(err)
				fmt.Fprint(w,string(b))
				return
			}
			tree,err = model.FindTreeById(id)
			if err != nil{
				resp.Status = "Fail"
				resp.Reason = "no id tree"
				b,_ := json.Marshal(resp)
				log.Println(err)
				fmt.Fprint(w,string(b))
				return
			}
		}

		if hash != ""{
			tree,err = model.FindTreeByHash(hash)
			if err != nil{
				resp.Status = "Fail"
				resp.Reason = "no hash tree"
				b,_ := json.Marshal(resp)
				log.Println(err)
				fmt.Fprint(w,string(b))
				return
			}

		}

		if tree == nil{
			resp.Status="Fail"
			resp.Reason="illegal query string"
			b,_ := json.Marshal(resp)
			fmt.Fprint(w,string(b))
			return
		}

		resp.Data = append(resp.Data,tree)
		log.Println("Querry ",tree)
		resp.Status="Success"
		b,_ := json.Marshal(resp)
		fmt.Fprint(w,string(b))
		return
	}
	if r.Method == "POST"{
		name := r.FormValue("name")
		height,_ := strconv.ParseFloat(r.FormValue("height"),32)
		hash := r.FormValue("hash")
		tree,_ := model.NewTree(name, hash, float32(height))
		err := tree.SaveOrUpdate()
		if err != nil{
			resp.Status = "Fail"
			b,_ := json.Marshal(resp)
			fmt.Fprint(w,string(b))
			return
		}
		resp.Status = "Success"
		b,_ := json.Marshal(resp)
		fmt.Fprint(w,string(b))
		return
	}
	if r.Method == "PUT"{
		r.ParseForm()
		id,_ := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		hash := r.FormValue("hash")
		height,_ := strconv.ParseFloat(r.FormValue("height"),32)
		tree := model.Tree{Id:id, Name:name, Hash:hash, Height:float32(height)}
		tree.SaveOrUpdate()

		resp.Status = "Success"
		b,_ := json.Marshal(resp)
		fmt.Fprint(w,string(b))
		return

	}
}

func GenHashHandler(w http.ResponseWriter, r *http.Request){
	var result HashResult
	identity:=r.FormValue("identity")
	name:=r.FormValue("name")
	hash := model.GenHash(identity,name)
	result.Status = "Success"
	result.Reason = ""
	result.Hash = hash
	b,_ := json.Marshal(result)
	fmt.Fprint(w,string(b))
	return
}