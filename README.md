# tdl

This tool is simple todos manager.

# Usage

```bash
# Add todo
$ tdl add "todo text"
# Show todo list
$ tdl show
# Delete todo in todo list
$ tdl delete 1
# Edit text of todo
$ tdl edit 2 "edited message"
# Star or unstar todo
$ tdl star 3
# Check or uncheck todo
$ tdl check 4
```

# Install And Setting

```bash
# Install this tool
$ go get -u github.com/x-color/tdl
# Setting todos file
$ mkdir ~/.tdl
$ echo "[]" > ~/.tdl/todos.json
```
