@echo off

:: Check if version argument is provided
if "%1"=="" (
    echo [ERROR] No version tag provided.
    echo Usage: build.bat ^<version-tag^>
    exit /b 1
)

:: Delete old build folder
rd /s /q ".\build\win"

:: Recreate build folder
mkdir build\win

:: Show argument
echo Building: sentrigo-%1

:: Build the Go binary
go build -o build\win\sentrigo-win-%1.exe cmd\sentrigo\main.go
