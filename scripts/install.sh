#!/bin/bash

PROGRAM_PATH=`pwd` 
PROGRAM_NAME=go-todo-list

sudo ln -sf $PROGRAM_PATH/$PROGRAM_NAME /bin/$PROGRAM_NAME

echo -e "\n\tSymbolic link created: /bin/$PROGRAM_NAME -> $PROGRAM_PATH/$PROGRAM_NAME"
