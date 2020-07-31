package main

import (
"database/sql"
"fmt"
_ "github.com/lib/pq"
"net/http"
"log"
"encoding/json"
"bytes"
"crypto/tls"
"os"
)

const (
host     = "172.28.198.148"
port     = 5432
user     = "nvxweb"
password = "TNd-6X"
dbname   = "web"
)

func main() {
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
"password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
db, err := sql.Open("postgres", psqlInfo)
if err != nil {
panic(err)
}

err = db.Ping()
if err != nil {
panic(err)
}
var q string
var arg string


 if (len(os.Args) > 1)  {
arg = os.Args[1]
}else{
arg =""	
}


if len(arg) > 0 {
q ="SELECT extension,password,user_context FROM v_extensions where extension is not null and user_context='"+arg+"'"
}else{
q ="SELECT extension,password,user_context FROM v_extensions where extension is not null"
}


rows, err := db.Query(q)	

if err != nil {
// handle this error better than this
panic(err)
return
}
defer rows.Close()

for rows.Next() {
var extension string
var password string
var user_context string
err = rows.Scan(&extension,&password,&user_context)
if err != nil {
// handle this error
panic(err)
}
user := extension+"@"+user_context
if len(user) > 0 {
message := map[string]interface{}{
"name": user,
"password":password,
"admin_channels":[]interface{}{
	""+user+"",
	"contact@"+user_context+"",
},
"admin_roles":[]interface{}{
	""+user+"" ,
},
"extension":extension,
"disabled":false,
"email":user,
}

bytesRepresentation, err := json.Marshal(message)
//log.Println(bytes.NewBuffer(bytesRepresentation))
if err != nil {
log.Fatalln(err)
}
http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
cbhost := "172.28.198.150"
cbvport := "4985"
cbbucket := "nvx"
url :=`https://` + cbhost + `:` + cbvport + `/` + cbbucket + `/_user/` +user+ ``
log.Println(url)

// initialize http client
client := &http.Client{}

// set the HTTP method, url, and request body
req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bytesRepresentation))
if err != nil {
panic(err)
}

// set the request header Content-Type for json
req.Header.Set("Content-Type", "application/json; charset=utf-8")
resp, err := client.Do(req)
if err != nil {
panic(err)
}

var result map[string]interface{}

json.NewDecoder(resp.Body).Decode(&result)

log.Println(resp.StatusCode)
//log.Println(result)

}else{
log.Println("user not found")
}

}

// get any error encountered during iteration
err = rows.Err()
if err != nil {
panic(err)
}

defer db.Close()


}


