# Todo Cli Application

This is a cli application written in go for managing todo's. [Project](https://roadmap.sh/projects/task-tracker)

## Tech Stack

- [Go](https://go.dev/)
- [Cobra cli](https://github.com/spf13/cobra)

## Installation

- First Clone the repo

```bash
git clone https://github.com/gideon-del/todo-cli.git
cd ./todo-cli
```

- Install the app

```bash
 go get
 go install
```

- Run to check

```bash
  todo
```

## Commands

1.  Add todo

```bash
todo add "your todo name"
```

2. Update todo by id

```bash
todo update 1 "your todo title"
```

3. Change todo status

```bash
todo mark-in-progress 1
todo mark-done 1
```

4. Delete todo

```bash
todo delete 1
```

5. List todo

```bash
todo list
```

6. Filter todo by status (done,in-progress,todo)

```bash
todo list done
```
