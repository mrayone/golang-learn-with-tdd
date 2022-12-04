install:
			go mod vendor && go mod tidy
checkErr:
			errcheck .