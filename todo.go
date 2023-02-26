package todo

type TodoList struct {
	Id          int `json:"id"`
	Title       int `json:"title"`
	Description int `json:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"bool"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
