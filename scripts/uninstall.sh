#!/bin/bash

PROGRAM_PATH=`pwd`
PROGRAM_NAME=go-todo-list

ls  $HOME/.go-todo-list &> /dev/null
if (( "$?" == 0 )); then
    rm -rf $HOME/.go-todo-list
    echo -e "\n\tFolder $HOME/.go-todo-list deleted"
else
    echo -e "\n\tFolder $HOME/.go-todo-list folder does not exist"
fi

ls  /bin/$PROGRAM_NAME &> /dev/null
if (( "$?" == 0 )); then
    echo -e "\nEnter root password to remove symlink /bin/$PROGRAM_NAME"
    sudo rm -f /bin/$PROGRAM_NAME

    if (( "$?" == 0 )); then
    echo -e "\n\tSymbolic link deleted: /bin/$PROGRAM_NAME -> $PROGRAM_PATH/$PROGRAM_NAME"
    else
        echo -e "\n\tFailed to delete symbolic link: /bin/$PROGRAM_NAME -> $PROGRAM_PATH/$PROGRAM_NAME"
    fi
else
    echo -e "\n\tSymbolic link does not exist: /bin/$PROGRAM_NAME -> $PROGRAM_PATH/$PROGRAM_NAME"
fi
