@echo off
echo Building ...
if exist random.exe del random.exe 
if exist main.exe del main.exe 
go build main.go 
ren main.exe random.exe
echo DONE !
