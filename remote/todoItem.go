package remote

type TodoItem struct {
	Id     string     `json:"_id"`
	Group  string     `json:"group"`
	Task   string     `json:"task"`
	Status TaskStatus `json:"status"`
}
type TaskStatus int64

const (
	Pending TaskStatus = iota
	InProgress
	Done
)

func NewTodoItem(group string, task string) *TodoItem {
	var item TodoItem
	item.Group = group
	item.Task = task
	item.Status = Pending

	return &item
}
