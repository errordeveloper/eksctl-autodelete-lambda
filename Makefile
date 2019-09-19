build:
	GOOS=linux GOARCH=amd64 go build -o ./list-all-clusters/eksctl-autodelete-lambda-list-all-clusters-func ./list-all-clusters