package main

type Repository struct {
	data map[string]string
}

type Service struct {
	repository Respository
}

type Server struct {
	service Service
}
