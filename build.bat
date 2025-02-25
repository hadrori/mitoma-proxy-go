@echo off
setlocal

echo Starting build...

if not exist bin (
    echo Creating bin/ ...
    mkdir bin
)

echo Resolve module dependencies...
go mod tidy

echo Build
go build -o bin\mitoma-proxy.exe -ldflags="-s -w" main.go

IF %ERRORLEVEL% NEQ 0 (
    echo Build Failed
    exit /b %ERRORLEVEL%
)

echo Build Completed
echo bin\mitoma-proxy.exe

endlocal