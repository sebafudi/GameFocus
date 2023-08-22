@echo off
setlocal
go build -ldflags "-H=windowsgui -s -w"
endlocal
