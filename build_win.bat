@if exist "%~dp0Release\cef2go.exe" (
    @del "%~dp0Release\cef2go.exe"
)

set GOPATH=%~dp0

go install cef
@if %ERRORLEVEL% neq 0 goto end

go install wingui
@if %ERRORLEVEL% neq 0 goto end

go build -o Release/cef2go.exe src/main_win.go
@if %ERRORLEVEL% neq 0 goto end

cd Release/
call "cef2go.exe"
cd ../

:end
@echo exit code = %ERRORLEVEL%
