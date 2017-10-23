# tree_guard api 文档

## Model

### Tree

```json

{"id":11,
  "name":"liushu",
  "hash":"sewrwergdsf",
  "height":12}

```
### User

```json

{"id":"1",
  "name":"admin",
  "Password":"27544304e5bfa887cad3eb1d61db66ce"}

```


## /gen_token

### GET

* params:
  - identity string 身份表示 (必选)(tmp)
  - name string 树种名字 (必选)
 
* result:
  - status string "Success"表示成功 "Fail"表示失败
  - Reason string 表示失败原因 成功空字串
  - hash string 生成的哈希字串
    
    ```json
    
    {"status":"Success","reason":"","hash":"aa4c9765a0fdc06e1a3a1b276f9bc1ec"}
    
    ```    

## /tree

### GET

* params:
  - id int (可选)
  - hash string (可选)
 
* result:
  - status string "Success"表示成功 "Fail"表示失败
  - Reason string 表示失败原因 成功空字串
  - data [Tree]list  一个Tree对象的list
      
    ```json
    {"status":"Success","data":[{"id":11,"name":"liushu","hash":"sewrwergdsf","height":12}],"reason":""}
    
    ```    
    
###  POST

* params:
  - hash
  - name
  - height
  
* result:
  - status string "Success"表示成功 "Fail"表示失败
  - Reason string 表示失败原因 成功空字串
  - data [Tree]list  一个Tree对象的list
    
  ```json
  {"status":"Success","data":[{"id":11,"name":"liushu","hash":"sewrwergdsf","height":12}],"reason":""}
  
  ```    
   
###  PUT

* params:
  - height
  - hash
  - name
  - height
  
* result:
  - status string "Success"表示成功 "Fail"表示失败
  - Reason string 表示失败原因 成功空字串
  - data [Tree]list  一个Tree对象的list
    
  ```json
  {"status":"Success","data":[{"id":11,"name":"liushu","hash":"sewrwergdsf","height":12}],"reason":""}
  
  ```  

## /user

### GET

* params:
  - id int (可选)
  - name string (可选)
 
* result:
  - status string "Success"表示成功 "Fail"表示失败
  - Reason string 表示失败原因 成功空字串
  - data [User]list  一个User对象的list
      
    ```json
    {"status":"Success","data":[{"id":"1","name":"admin","Password":"27544304e5bfa887cad3eb1d61db66ce"}],"reason":""}
    
    ```    

###  POST

* params:
  - password
  - name
 
  
* result:
  - status string "Success"表示成功 "Fail"表示失败
  - Reason str  ing 表示失败原因 成功空字串
  - data [User]list  一个User对象的list
    
  ```json
  {"status":"Success","data":[{"id":"1","name":"admin","Password":"27544304e5bfa887cad3eb1d61db66ce"}],"reason":""}
  
  ```    
 
## /login

### GET

获取登陆表单 一般不用

### POST


* params:
  - password
  - username
 
  
* result:
  - status string "Success"表示成功 "Fail"表示失败
  - Reason string 表示失败原因 成功空字串
  - data [User]list  一个Tree对象的list
    
  ```json
  {"status":"Success","data":[{"id":"1","name":"admin","Password":"27544304e5bfa887cad3eb1d61db66ce"}],"reason":""}
  
  ```    
  
## /logout

### GET

无参数 根据cookie

## /user_test

### GET

测试user登陆情况
