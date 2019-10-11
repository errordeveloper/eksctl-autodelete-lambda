funcs := list-clusters

build:
	for func in $(funcs) ; do env GOARCH=amd64 GOOS=linux go build -o ./$${func}/eksctl-autodelete-lambda-$${func}-func ./$${func} ; done