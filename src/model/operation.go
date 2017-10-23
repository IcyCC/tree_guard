package model

import (
	"time"
	"errors"
	"log"
)

const (
	ADD_TREE = "ADD_TREE"
	UPDATE_TREE = "UPDATE_TREE"
	LOOK_UP_TREE = "LOOK_UP_TREE"
)

type Operation struct {
	Id        int
	Operate   string
	Timestamp time.Time
	User      *User
	Tree      *Tree
}

func FindOperationById(id int) (*Operation, error){

	var op Operation
	err := DB.QueryRow(`SELECT r.id,r.operation,r.user_id,r.user_name,r.timestamp,t.hash,t.name,t.height
	FROM tree AS t RIGHT JOIN
  (SELECT op.id as id,op.operation as operation,
    op.timestamp as timestamp, op.user_id as user_id,
    u.name as user_name, op.tree_hash as tree_hash
   FROM operation AS op JOIN user
     AS u on op.user_id=u.id WHERE op.id = ? ) AS r ON r.tree_hash = t.hash`,id).Scan(&op.Id, &op.Operate,&op.User.Id, &op.User.Name, &op.Timestamp, &op.Tree.Hash,&op.Tree.Name,op.Tree.Height)

	if err != nil{
		return nil,errors.New("no result")
	}

	return &op, nil
}

func FindOperationByName(operation string) (*Operation, error){

	var op Operation
	err := DB.QueryRow(`SELECT r.id,r.operation,r.user_id,r.user_name,r.timestamp,t.hash,t.name,t.height
	FROM tree AS t RIGHT JOIN
  (SELECT op.id as id,op.operation as operation,
    op.timestamp as timestamp, op.user_id as user_id,
    u.name as user_name, op.tree_hash as tree_hash
   FROM operation AS op JOIN user
     AS u on op.user_id=u.id WHERE op.operation = ? ) AS r ON r.tree_hash = t.hash`,operation).Scan(&op.Id, &op.Operate,&op.User.Id, &op.User.Name, &op.Timestamp, &op.Tree.Hash,&op.Tree.Name,op.Tree.Height)

	if err != nil{
		return nil,errors.New("no result")
	}

	return &op, nil
}

func FindOperationByUserId(userId int)(*Operation, error){

	var op Operation
	err := DB.QueryRow(`SELECT r.id,r.operation,r.user_id,r.user_name,r.timestamp,t.hash,t.name,t.height
FROM tree AS t RIGHT JOIN
  (SELECT op.id as id,op.operation as operation,
          op.timestamp as timestamp, op.user_id as user_id,
          u.name as user_name, op.tree_hash as tree_hash
   FROM operation AS op JOIN user
     AS u on op.user_id=u.id WHERE u.id = ? ) AS r ON r.tree_hash = t.hash`,userId).Scan(&op.Id, &op.Operate,&op.User.Id, &op.User.Name, &op.Timestamp, &op.Tree.Hash,&op.Tree.Name,op.Tree.Height)

	if err != nil{
		return nil,errors.New("no result")
	}

	return &op, nil
}

func FindOperationByTreeHash(hash string)(*Operation, error){

	var op Operation
	err := DB.QueryRow(`SELECT r.id,r.operation,r.user_id,r.user_name,r.timestamp,t.hash,t.name,t.height
FROM tree AS t RIGHT JOIN
  (SELECT op.id as id,op.operation as operation,
          op.timestamp as timestamp, op.user_id as user_id,
          u.name as user_name, op.tree_hash as tree_hash
   FROM operation AS op JOIN user
     AS u on op.user_id=u.id WHERE t.hash = ? ) AS r ON r.tree_hash = t.hash`,hash).Scan(&op.Id, &op.Operate,&op.User.Id, &op.User.Name, &op.Timestamp, &op.Tree.Hash,&op.Tree.Name,op.Tree.Height)

	if err != nil{
		return nil,errors.New("no result")
	}

	return &op, nil
}

func (op *Operation)SaveUpdate()error{
		res,err := DB.Exec(`INSERT INTO operation(operation,user_id,tree_hash,timestamp) VALUES (?,?,?,?) ON
					DUPLICATE KEY UPDATE operation = VALUES(operation),user_id=VALUES(user_id),
					tree_hash = VALUES(tree_hash),timestamp = VALUES(timestamp)`,
						op.User.Id,op.User.Name, op.Tree.Hash,op.Timestamp)
	if err!=nil{
		log.Println("Tree id : ", op.Id," save fail")
		return err
	}
	log.Println("insert success eff id : ", res)
	return nil
}

func NewOperaion(op string,u *User,t *Tree)*Operation{
	return &Operation{Operate:op,Timestamp:time.Now(),User:u,Tree:t}
}

