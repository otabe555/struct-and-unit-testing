package tasklist

type Task struct {
	ID int
	Assignee,
	Title,
	Deadline string
	Done bool
}

type TaskList struct {
	list []Task
}

func NewTasklist() *TaskList {
	p := new(TaskList)
	p.list = make([]Task, 0)
	return p
}

func (that *TaskList) Create(task *Task) *Task {
	for _, element := range that.list {
		if element.ID == task.ID {
			return &Task{}
		}
	}
	that.list = append(that.list, *task)
	return task
}
func (that *TaskList) Update(task *Task) *Task {
	for index, element := range that.list {
		if element.ID == task.ID {
			that.list[index] = *task
			return &that.list[index]
		}
	}
	return &Task{}
}
func (that *TaskList) Delete(ID int) *Task {
	for index, element := range that.list {
		if element.ID == ID {
			that.list = append(that.list[:index], that.list[index+1:]...)
			return &element
		}
	}
	return &Task{}
}
func (that *TaskList) MakeDone(id int) *Task {
	for index, element := range that.list {
		if element.ID == id {
			that.list[index].Done = true
			return &that.list[index]
		}
	}
	return &Task{}
}
func (that *TaskList) Get(id int) *Task {
	for index, element := range that.list {
		if element.ID == id {
			return &that.list[index]
		}
	}
	return &Task{}
}
func (that *TaskList) GetAll() *[]Task {
	return &that.list
}
