package main

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

func CreateNewTasklist() *TaskList {
	p := new(TaskList)
	p.list = make([]Task, 0)
	return p

}
func (this *TaskList) Create(task *Task) *Task {
	for _, element := range this.list {
		if element.ID == task.ID {
			return &Task{}
		}
	}
	this.list = append(this.list, *task)
	return task
}
func (this *TaskList) Update(task *Task) *Task {
	for _, element := range this.list {
		if element.ID == task.ID {
			element = *task
			return task
		}
	}
	return &Task{}
}
func (this *TaskList) Delete(ID int) *Task {
	for index, element := range this.list {
		if element.ID == ID {
			this.list = append(this.list[:index], this.list[index+1:]...)
			return &element
		}
	}
	return &Task{}
}
func (this *TaskList) MakeDone(task *Task) *Task {
	for _, element := range this.list {
		if element.ID == task.ID {
			element.Done = true
			return task
		}
	}
	return &Task{}
}
func (this *TaskList) Get(task *Task) *Task {
	for _, element := range this.list {
		if element.ID == task.ID {
			return task
		}
	}
	return &Task{}
}
func (this *TaskList) GetAll() *[]Task {
	return &this.list
}
