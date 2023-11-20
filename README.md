# GoTodoList

This app was created just for learning Golang ¯\\\_(ツ)_/¯.


## Build app

```bash
make
```

After building you get an executable file `go-todo-list`.


## Install app

```bash
make install
```

After this step will be created symlink `/bin/go-todo-list`.


## Uninstall app

```bash
make uninstall
```

Remove the symlink and `~/.go-todo-list folder` which is created after the first start of the program.


## Usage

### Insert new task into TODO-list

Create a new task in the `SQLite` database. The database file will be stored in `~/.go-todo-list/go-todo-list.db`

```bash
go-todo-list add -n "My first task" -d "Simple description"
```

As output, you will receive a task ID in the database.

```bash
Task created with ID: 1
```

### Remove the task from TODO-list by ID

```bash
go-todo-list delete -i 1
```

As output, you will receive a notification that deletion was performed successfully (or error 👉👈).

```bash
Task with ID (1) was deleted
```


### Mark the task as completed

```bash
go-todo-list done -i 2
```

As in the previous command, you will be informed about the execution status.

```bash
Task was successfully updated
```


### Show all task

```bash
go-todo-list show
```

The program will print task info as a table.

```bash
ID  NAME           DESCRIPTION         STATUS  
2   My first task  Simple description  DONE    
3   Hello)         Simple description  NEW 
```
