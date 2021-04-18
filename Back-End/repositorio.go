package main

func NewService(repository Repository) Service {
	return Service{repository}
}
