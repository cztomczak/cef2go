@if exist "%~dp0bin\cef2go.exe" (
    @del "%~dp0bin\cef2go.exe"
)
go build -o bin/cef2go.exe src/main_win.go
call "bin/cef2go.exe"
