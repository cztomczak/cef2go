@if exist "%~dp0Release\cef2go.exe" (
    @del "%~dp0Release\cef2go.exe"
)

set GOPATH=%~dp0

go install cef
@if %ERRORLEVEL% neq 0 goto end

go install wingui
@if %ERRORLEVEL% neq 0 goto end

@SET ORIGPATH=%PATH%
@SET PATH=%PATH%;%~dp0/Release
go test src/tests/cef_test.go
@if %ERRORLEVEL% neq 0 goto end
@SET PATH=%ORIGPATH%
@if exist "%~dp0src\tests\debug.log" (
    @del "%~dp0src\tests\debug.log"
)

go build -o Release/cef2go.exe src/main_windows.go
@if %ERRORLEVEL% neq 0 goto end

cd Release/
call "cef2go.exe"
cd ../

:end
@echo exit code = %ERRORLEVEL%
