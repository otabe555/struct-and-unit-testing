package tasklist

import "testing"

var (
	NewTask = &Task{
		ID:       5,
		Assignee: "John",
		Title:    "Do homework",
		Deadline: "5th march",
		Done:     false,
	}
	taskList = NewTasklist()
)

func TestCreate(t *testing.T) {
	createdTask := taskList.Create(NewTask)
	if createdTask != NewTask {
		t.Error("Fail to create")
	}
}
func TestUpdate(t *testing.T) {
	UpdatingTask := &Task{
		ID:       5,
		Assignee: "Mars",
		Title:    "Do homework",
		Deadline: "5th march",
		Done:     true,
	}
	createdTask := taskList.Update(UpdatingTask)
	if *createdTask != *UpdatingTask {
		t.Error("Fail to Update", createdTask, UpdatingTask)
	}
}
func TestDelete(t *testing.T) {
	taskList.Create(NewTask)
	if taskList.Delete(5).ID != 5 {
		t.Error("Failed to delete")
	}

}
func TestMakeDone(t *testing.T) {
	taskList.Create(NewTask)

	if !taskList.MakeDone(5).Done {
		t.Error("Failed to set it done")
	}
}
func TestGet(t *testing.T) {
	createdTask := taskList.Create(NewTask)
	if *taskList.Get(createdTask.ID) != *createdTask {
		t.Error("Failed to get the Task specifications")
	}
}

func TestGetAll(t *testing.T) {
	if &taskList.list != taskList.GetAll() {
		t.Error("Couldnt get the full list of tasks`")
	}
}
