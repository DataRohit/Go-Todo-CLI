package todo

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type TodoItems []Todo

func (t *TodoItems) Add(title string) {
	*t = append(*t, Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	})
}

func (t TodoItems) ValidateIndex(index int) error {
	if index < 0 || index >= len(t) {
		return errors.New("invalid index")
	}
	return nil
}

func (t *TodoItems) Delete(index int) error {
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
	return nil
}

func (t *TodoItems) Toggle(index int) error {
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	(*t)[index].Completed = !(*t)[index].Completed
	if (*t)[index].Completed {
		now := time.Now()
		(*t)[index].CompletedAt = &now
	} else {
		(*t)[index].CompletedAt = nil
	}
	return nil
}

func (t *TodoItems) Edit(index int, title string) error {
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	(*t)[index].Title = title
	return nil
}

func (t TodoItems) List() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, item := range t {
		completed := "❌"
		completedAt := ""
		if item.Completed {
			completed = "✅"
			completedAt = item.CompletedAt.Format(time.RFC1123)
		}
		table.AddRow(strconv.Itoa(index), item.Title, completed, item.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
