package main

import "fmt"

// Task's list
type TaskList struct {
	//Array de tipo TASK
	tasks []*Task
}

func (tl *TaskList) addTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(i int) {
	tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
}

// =========================================================================================================

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) toPrint() {
	fmt.Printf("Name: %s\nDescription: %s\nCompleted: %t\n", t.name, t.description, t.completed)
	fmt.Println("=======================================")

}

func (t *Task) setCompleted() {
	t.completed = true
}

func main() {
	t1 := Task{
		name:        "Curso de GO",
		description: "Completar curso de go este mes",
	}

	t2 := Task{
		name:        "Curso HTML",
		description: "Completar esta Semana",
	}

	list := TaskList{}
	list.addTask(&t1)
	list.addTask(&t2)

	// t1.toPrint()
	// t2.toPrint()
	fmt.Println(list)

	t3 := Task{
		name:        "Curso CSS",
		description: "Completar esta Semana",
	}

	list.addTask(&t3)

	list.removeTask(1)

	for _, v := range list.tasks {
		fmt.Println(v.name)
	}

}
