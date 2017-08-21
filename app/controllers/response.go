package controllers

type JsonResponse struct {
	Success bool   `json:"success"`
	Html    string `json:"html"`
	Message string `json:"message"`
}
