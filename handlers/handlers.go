package handlers

import "net/http"

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}