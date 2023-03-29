package main

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Todo `json:"data"`
	Message string `json:"message"`
}
