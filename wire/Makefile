# git commit message
msg ?=
# git branch
b = main

run:
	wire
	go build -o main .
	./main

push:
ifeq ($(strip $(msg)),)
	@echo "Need commit message with flag 'make push msg='"
else
	@echo ${msg}
endif
	go fix
	git add .
	git commit -m "${msg}"
	@echo "push branch to origin/${b}"
	git push origin ${b}