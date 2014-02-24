set GOPATH=%~dp0

go install cef
@if %ERRORLEVEL% neq 0 goto end

go install wingui
@if %ERRORLEVEL% neq 0 goto end

IF "%1"=="http_server" (
    @if exist "%~dp0Release\cef2go_http_server.exe" (
        @del "%~dp0Release\cef2go_http_server.exe"
    )
    go build -o Release/cef2go_http_server.exe src/http_server_windows.go
    @if %ERRORLEVEL% neq 0 goto end
    cd Release/
    call "cef2go_http_server.exe"
    cd ../
    goto end
)

@SET ORIGPATH=%PATH%
@SET PATH=%PATH%;%~dp0/Release
go test src/tests/cef_test.go
@if %ERRORLEVEL% neq 0 goto end
@SET PATH=%ORIGPATH%
@if exist "%~dp0src\tests\debug.log" (
    @del "%~dp0src\tests\debug.log"
)

@if exist "%~dp0Release\cef2go.exe" (
    @del "%~dp0Release\cef2go.exe"
)

IF "%1"=="noconsole" (
    go build -ldflags="-H windowsgui" -o Release/cef2go.exe src/main_windows.go
    @if %ERRORLEVEL% neq 0 goto end
) else (
    go build -o Release/cef2go.exe src/main_windows.go
    @if %ERRORLEVEL% neq 0 goto end
)

cd Release/
call "cef2go.exe"
cd ../

:end
@echo exit code = %ERRORLEVEL%
