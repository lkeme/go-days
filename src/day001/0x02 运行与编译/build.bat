@echo off
echo Go build running

set packageName="cmd"
set appName="Hello"
set buildVersion="v1.0.4"

cd %packageName%

set GOOS=windows
set GOARCH=amd64
go build -o ../build/%appName%_%GOOS%_%GOARCH%_%buildVersion%.exe
echo Done %appName%_%GOOS%_%GOARCH%_%buildVersion%

set GOOS=linux
set GOARCH=amd64
go build -o ../build/%appName%_%GOOS%_%GOARCH%_%buildVersion%
echo Done %appName%_%GOOS%_%GOARCH%_%buildVersion%

set GOOS=linux
set GOARCH=arm64
go build -o ../build/%appName%_%GOOS%_%GOARCH%_%buildVersion%
echo Done %appName%_%GOOS%_%GOARCH%_%buildVersion%

set GOOS=darwin
set GOARCH=amd64
go build -o ../build/%appName%_%GOOS%_%GOARCH%_%buildVersion%
echo Done %appName%_%GOOS%_%GOARCH%_%buildVersion%

pause