package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("invalid index")
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	err := todos.validateIndex(index)

	if err != nil {
		return err
	}

	t := *todos
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	err := todos.validateIndex(index)

	if err != nil {
		return err
	}

	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) toggle(index int) error {
	err := todos.validateIndex(index)

	if err != nil {
		return err
	}

	isCompleted := (*todos)[index].Completed
	if !isCompleted {
		(*todos)[index].Completed = true

		now := time.Now()
		(*todos)[index].CompletedAt = &now
	}

	return fmt.Errorf("this todo is already completed at: %v", (*todos)[index].CompletedAt)
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")

	for index, t := range *todos {
		completed := "❌"
		completedAt := "-"

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
