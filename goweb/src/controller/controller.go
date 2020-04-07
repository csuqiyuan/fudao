package controller

import (
	"log"
	"net/http"
	"wserver/src/service"
)


func Init(){
	getSubject()
	getGrades()
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getSubject(){
	http.HandleFunc("/get_subject", service.GetSubject) //设置访问的路由
}

func getGrades(){
	http.HandleFunc("/get_grades", service.GetGrades)
}