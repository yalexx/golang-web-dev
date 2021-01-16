package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Code struct {
	Code    int    `json:"code"`
	Descrip string `json:"descrip"`
}

type Codes []Code

func main() {
	var data Codes
	rcvd := `[{"code":200,"descrip":"StatusOK"},{"code":301,"descrip":"StatusMovedPermanently"},{"code":302,"descrip":"StatusFound"},{"code":303,"descrip":"StatusSeeOther"},{"code":307,"descrip":"StatusTemporaryRedirect"},{"code":400,"descrip":"StatusBadRequest"},{"code":401,"descrip":"StatusUnauthorized"},{"code":402,"descrip":"StatusPaymentRequired"},{"code":403,"descrip":"StatusForbidden"},{"code":404,"descrip":"StatusNotFound"},{"code":405,"descrip":"StatusMethodNotAllowed"},{"code":418,"descrip":"StatusTeapot"},{"code":500,"descrip":"StatusInternalServerError"}]`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
	for _, v := range data {
		fmt.Println(v.Code)
		fmt.Println(v.Descrip)
	}
}
