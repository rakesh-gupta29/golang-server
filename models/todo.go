package models

type Todo struct {
	ID        string `json:"id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	Desc      string `json:"desc" bson:"desc"`
	Date      string `json:"date" bson:"date"`
	Completed bool   `json:"completed" bson:"completed"`
}
