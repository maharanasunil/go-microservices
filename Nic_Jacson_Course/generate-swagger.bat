@echo off
where swagger >nul 2>nul
if %errorlevel% neq 0 (
    echo Swagger CLI not found. Installing...
    go install github.com/go-swagger/go-swagger/cmd/swagger@latest
) else (
    echo Swagger already installed.
)

echo Generating swagger.yaml ...
swagger generate spec -o ./swagger.yaml --scan-models
echo Swagger spec generated successfully!
pause
