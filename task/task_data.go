package task

type TaskData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewTaskData(title, description string) *TaskData {
	return &TaskData{
		Title:       title,
		Description: description,
	}
}
