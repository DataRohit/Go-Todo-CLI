# Go Todo CLI

```bash
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go --help
Usage of /tmp/go-build2426606230/b001/exe/main:
  -add string
        Add a new todo by specifying title
  -del int
        Delete a todo by specifying its index (default -1)
  -edit string
        Edit a todo by index & specify a new title (format: id:new_title)
  -list
        List all todoItems
  -toggle int
        Toggle a todo's completion status by index (default -1)
```

```bash
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -add "Create a TODO CLI using Go Lang"
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -list
┌───┬─────────────────────────────────┬───────────┬───────────────────────────────┬──────────────┐
│ # │              Title              │ Completed │          Created At           │ Completed At │
├───┼─────────────────────────────────┼───────────┼───────────────────────────────┼──────────────┤
│ 0 │ Create a TODO CLI using Go Lang │ ❌        │ Sun, 29 Sep 2024 13:39:26 IST │              │
└───┴─────────────────────────────────┴───────────┴───────────────────────────────┴──────────────┘
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -add "Add Edit & Delete Functionality"
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -add "Add Toggle & List Functionality"
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -list
┌───┬─────────────────────────────────┬───────────┬───────────────────────────────┬──────────────┐
│ # │              Title              │ Completed │          Created At           │ Completed At │
├───┼─────────────────────────────────┼───────────┼───────────────────────────────┼──────────────┤
│ 0 │ Create a TODO CLI using Go Lang │ ❌        │ Sun, 29 Sep 2024 13:39:26 IST │              │
│ 1 │ Add Edit & Delete Functionality │ ❌        │ Sun, 29 Sep 2024 13:40:02 IST │              │
│ 2 │ Add Toggle & List Functionality │ ❌        │ Sun, 29 Sep 2024 13:40:17 IST │              │
└───┴─────────────────────────────────┴───────────┴───────────────────────────────┴──────────────┘
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -toggle 0
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -list
┌───┬─────────────────────────────────┬───────────┬───────────────────────────────┬───────────────────────────────┐
│ # │              Title              │ Completed │          Created At           │         Completed At          │
├───┼─────────────────────────────────┼───────────┼───────────────────────────────┼───────────────────────────────┤
│ 0 │ Create a TODO CLI using Go Lang │ ✅        │ Sun, 29 Sep 2024 13:39:26 IST │ Sun, 29 Sep 2024 13:40:29 IST │
│ 1 │ Add Edit & Delete Functionality │ ❌        │ Sun, 29 Sep 2024 13:40:02 IST │                               │
│ 2 │ Add Toggle & List Functionality │ ❌        │ Sun, 29 Sep 2024 13:40:17 IST │                               │
└───┴─────────────────────────────────┴───────────┴───────────────────────────────┴───────────────────────────────┘
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -toggle 1
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -list
┌───┬─────────────────────────────────┬───────────┬───────────────────────────────┬───────────────────────────────┐
│ # │              Title              │ Completed │          Created At           │         Completed At          │
├───┼─────────────────────────────────┼───────────┼───────────────────────────────┼───────────────────────────────┤
│ 0 │ Create a TODO CLI using Go Lang │ ✅        │ Sun, 29 Sep 2024 13:39:26 IST │ Sun, 29 Sep 2024 13:40:29 IST │
│ 1 │ Add Edit & Delete Functionality │ ✅        │ Sun, 29 Sep 2024 13:40:02 IST │ Sun, 29 Sep 2024 13:40:44 IST │
│ 2 │ Add Toggle & List Functionality │ ❌        │ Sun, 29 Sep 2024 13:40:17 IST │                               │
└───┴─────────────────────────────────┴───────────┴───────────────────────────────┴───────────────────────────────┘
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -del 1
(base) datarohit@DataRohit:~/Desktop/Go-ToDo-CLI$ go run main.go -list
┌───┬─────────────────────────────────┬───────────┬───────────────────────────────┬───────────────────────────────┐
│ # │              Title              │ Completed │          Created At           │         Completed At          │
├───┼─────────────────────────────────┼───────────┼───────────────────────────────┼───────────────────────────────┤
│ 0 │ Create a TODO CLI using Go Lang │ ✅        │ Sun, 29 Sep 2024 13:39:26 IST │ Sun, 29 Sep 2024 13:40:29 IST │
│ 1 │ Add Toggle & List Functionality │ ❌        │ Sun, 29 Sep 2024 13:40:17 IST │                               │
└───┴─────────────────────────────────┴───────────┴───────────────────────────────┴───────────────────────────────┘
```
