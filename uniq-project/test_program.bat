@echo off
echo Building program...
go build -o uniq.exe cmd/uniq/main.go

echo.
echo === Test 1: Basic functionality ===
echo hello > test1.txt
echo world >> test1.txt
echo hello >> test1.txt
echo test >> test1.txt
echo world >> test1.txt

echo Input:
type test1.txt

echo.
echo Output without flags:
.\uniq.exe test1.txt

echo.
echo Output with -c flag:
.\uniq.exe -c test1.txt

echo.
echo Output with -d flag:
.\uniq.exe -d test1.txt

echo.
echo Output with -u flag:
.\uniq.exe -u test1.txt

del test1.txt