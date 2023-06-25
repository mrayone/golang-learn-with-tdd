install:
			go mod vendor && go mod tidy
checkErr:
			errcheck .

test:
	go test -race ./... -count=1 -short