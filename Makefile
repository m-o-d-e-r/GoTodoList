PROGRAM_NAME=go-todo-list

build:
	go build -o ${PROGRAM_NAME}

install:
	scripts/install.sh 

uninstall:
	scripts/uninstall.sh 
