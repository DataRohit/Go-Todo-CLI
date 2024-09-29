package todo

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo by specifying title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title (format: id:new_title)")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by specifying its index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo's completion status by index")
	flag.BoolVar(&cf.List, "list", false, "List all todoItems")
	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todoItems *TodoItems) {
	switch {
	case cf.List:
		todoItems.List()

	case cf.Add != "":
		todoItems.Add(cf.Add)

	case cf.Edit != "":
		if index, title, err := parseEditFlag(cf.Edit); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		} else {
			todoItems.Edit(index, title)
		}

	case cf.Toggle != -1:
		todoItems.Toggle(cf.Toggle)

	case cf.Del != -1:
		todoItems.Delete(cf.Del)

	default:
		fmt.Println("Invalid command. Use --help to see available options.")
	}
}

func parseEditFlag(editFlag string) (int, string, error) {
	parts := strings.SplitN(editFlag, ":", 2)
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("invalid format for edit, use id:new_title")
	}
	index, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", fmt.Errorf("invalid index for edit")
	}
	return index, parts[1], nil
}
