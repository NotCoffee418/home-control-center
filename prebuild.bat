@echo off
echo Building frontend...
cd frontend

if not exist node_modules (
    echo Installing npm dependencies...
    npm install
)

echo Running npm run build...
npm run build

if not exist dist (
    echo ERROR: dist directory not created
    exit /b 1
)

if not exist dist\index.html (
    echo ERROR: index.html not found in dist
    exit /b 1
)

echo Frontend build successful
dir dist