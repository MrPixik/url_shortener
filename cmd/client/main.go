package main

import (
	"github.com/go-resty/resty/v2"
)

/*
Server testing manually
*/
//TODO	Добавить считывание ошибок и ответов сервера в отдельные easyjson-объекты
//		Создать меню, с выбором куда отправить запрос
//		Узнать можно ли использовать easyjson с resty
func main() {
	client := resty.New()

	client.
		SetBaseURL("http://localhost:8080/").
		SetRedirectPolicy(resty.NoRedirectPolicy())

	menu(client)
	//var reqUrl requestUrl
	//mainPageGetRequest(client, reqUrl)
}
