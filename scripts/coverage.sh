mkdir -p .coverage
go test . -coverprofile=.coverage/coverage.txt -covermode=atomic
go tool cover -html=.coverage/coverage.txt
