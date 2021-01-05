/*
 * Copyright (C) 2021 Baidu, Inc. All Rights Reserved.
 */
package pattern

import (
	"fmt"
	"log"
	"net/http"
)

//func HomePage(w http.ResponseWriter ,r *http.Request)  {
//	fmt.Println("Endpoint Hit: homePage")
//	fmt.Fprintf(w, "Welcome to the HomePage!")
//}
//
//func HandleRequests()  {
//	http.HandleFunc("/",HomePage)
//	log.Fatal(http.ListenAndServe(":8081", nil))
//}
//func TestDecorator() {
//	HandleRequests()
//}

//装饰器方法，用于增强HomePage的功能，此例子加入了http访问简单的权限校验
func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking to see if Authorized header set...")
		if val,ok := r.Header["Authorized"]; ok {
			if val[0] =="true" {
				fmt.Println("Header is set! We can serve content!")
				endpoint(w,r)
			}
		} else {
			fmt.Println("Not Authorized!!")
			fmt.Fprintf(w, "Not Authorized!!")
		}
	})
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func HandleRequests() {
	http.Handle("/", isAuthorized(HomePage))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func TestDecorator() {
	HandleRequests()
}
