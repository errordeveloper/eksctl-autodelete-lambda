module github.com/errordeveloper/eksctl-autodelete-lambda

go 1.12

require (
	github.com/aws/aws-lambda-go v1.13.2

	github.com/weaveworks/eksctl v0.0.0-20191028162209-a02037f03d2b
)

replace (
	// Override since git.apache.org is down.  The docs say to fetch from github.
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v10.14.0+incompatible
	// Needed due to to Sirupsen/sirupsen case clash
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/awslabs/goformation => github.com/errordeveloper/goformation v0.0.0-20190507151947-a31eae35e596
	// go mod appears to pick wrong version of github.com/docker/distribution automatically, which breaks k8s.io/kubernetes@v1.12.6
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20190619192407-5223c27422cc
	// Used to pin the k8s library versions regardless of what other dependencies enforce
	k8s.io/api => k8s.io/api v0.0.0-20190226173710-145d52631d00
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190226180157-bd0469a053ff
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190226174127-78295b709ec6
)
