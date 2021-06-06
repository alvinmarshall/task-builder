package dto

type TaskRequest struct {
	Title       string `form:"title"`
	IsCompleted bool   `form:"is_completed"`
}
type TaskResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}
