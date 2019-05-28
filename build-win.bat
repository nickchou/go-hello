#CGO_ENABLED=0 最新版本已经不需要这个了
#GOOS=linux GOARCH=amd64 go build main.go

set GOOS=linux
set GOARCH=amd64
go build main.go

pause