@echo off
setlocal

set BINARY_NAME=golang-server
set BUILD_DIR=build
set MAIN_PATH=./cmd/server

if "%1"=="build" goto build
if "%1"=="run" goto run
if "%1"=="test" goto test
if "%1"=="test-coverage" goto test-coverage
if "%1"=="clean" goto clean
if "%1"=="lint" goto lint
if "%1"=="format" goto format
if "%1"=="deps" goto deps
if "%1"=="docs" goto docs
if "%1"=="docker-build" goto docker-build
if "%1"=="docker-run" goto docker-run
if "%1"=="install-tools" goto install-tools
if "%1"=="help" goto help

echo Unknown command: %1
echo Use 'make.bat help' to see available commands
exit /b 1

:build
echo Building application...
if not exist %BUILD_DIR% mkdir %BUILD_DIR%
go build -o %BUILD_DIR%\%BINARY_NAME%.exe %MAIN_PATH%
goto end

:run
echo Running application...
go run %MAIN_PATH%
goto end

:test
echo Running tests...
go test -v ./...
goto end

:test-coverage
echo Running tests with coverage...
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
echo Coverage report generated: coverage.html
goto end

:clean
echo Cleaning build files...
if exist %BUILD_DIR% rmdir /s /q %BUILD_DIR%
if exist coverage.out del coverage.out
if exist coverage.html del coverage.html
goto end

:lint
echo Running linter...
golangci-lint run
goto end

:format
echo Formatting code...
go fmt ./...
go vet ./...
goto end

:deps
echo Installing dependencies...
go mod tidy
go mod download
goto end

:docs
echo Generating documentation...
godoc -http=:6060
goto end

:docker-build
echo Building Docker image...
docker build -t %BINARY_NAME% .
goto end

:docker-run
echo Running Docker container...
docker run -p 8080:8080 %BINARY_NAME%
goto end

:install-tools
echo Installing development tools...
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/godoc@latest
goto end

:help
echo Available commands:
echo   build         - Build the application
echo   run           - Run the application
echo   test          - Run tests
echo   test-coverage - Run tests with coverage report
echo   clean         - Clean build files
echo   lint          - Run linter
echo   format        - Format code
echo   deps          - Install dependencies
echo   docs          - Generate documentation
echo   docker-build  - Build Docker image
echo   docker-run    - Run Docker container
echo   install-tools - Install development tools
echo   help          - Show this help
goto end

:end 