package model

type AboutItemModel struct {
	ID 		int
	Title 	string
	Content	string
}

type AboutResponseModel struct {
	Data	[]AboutItemModel
	Total	int
}