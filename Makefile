#
# Makefile for golang-example
#
.PHONY: usage edit build clean git
#----------------------------------------------------------------------------------
usage:
	@echo "make [edit|build]"
#----------------------------------------------------------------------------------
edit e:
	@echo "make (edit:e) [history]"
edit-go eg:
	vi main.go
edit-history eh:
	vi HISTORY.md
#----------------------------------------------------------------------------------
build b:
#----------------------------------------------------------------------------------
clean:
#----------------------------------------------------------------------------------
run r:
	@echo "make (run:r) [hook|chat]"
#----------------------------------------------------------------------------------
git g:
	@echo "make (git:g) [update|store]"
git-update gu:
	git add .
	git commit -a -m "update contents"
	git push
git-store gs:
	git config credential.helper store
#----------------------------------------------------------------------------------

