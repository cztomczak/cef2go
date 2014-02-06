@if exist "%~dp0Release\cef2go.exe" (
    @del "%~dp0Release\cef2go.exe"
)
go build -o Release/cef2go.exe src/main_win.go
cd Release/
call "cef2go.exe"
cd ../
