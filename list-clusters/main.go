package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
	"github.com/weaveworks/eksctl/pkg/eks"
)

func doListCluster(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	region, ok := request.QueryStringParameters["region"]
	if !ok {
		region = api.DefaultRegion
	}

	minAgeString, ok := request.QueryStringParameters["minAge"]
	if !ok {
		minAgeString = "0s"
	}
	minAge, err := time.ParseDuration(minAgeString)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	cfg := &api.ClusterConfig{
		Metadata: &api.ClusterMeta{
			Region: region,
			Tags:   map[string]string{},
		},
	}

	ctl := eks.New(&api.ProviderConfig{Region: region}, cfg)

	if err := ctl.CheckAuth(); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 502}, err
	}

	stackManager := ctl.NewStackManager(cfg)

	// TODO: convert this to ListOwnedStacksByClusterName
	stacks, err := stackManager.ListStacksMatching(".*")
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 502}, err
	}

	type stackDescriptor struct {
		CreatedAt   time.Time
		Creator, ID string
	}

	stacksByName := map[string][]stackDescriptor{}

	for _, s := range stacks {
		for _, tag := range s.Tags {
			if *tag.Key == api.ClusterNameTag {
				if time.Since(*s.CreationTime) < minAge {
					continue
				}

				stack := stackDescriptor{
					ID:        *s.StackId,
					CreatedAt: *s.CreationTime,
				}

				stackEvents, err := stackManager.LookupCloudTrailEvents(s)
				if err != nil {
					return events.APIGatewayProxyResponse{StatusCode: 502}, err
				}
				stack.Creator = fmt.Sprintf("%v", stackEvents)
				for _, e := range stackEvents {
					if *e.EventSource != "cloudformation.amazonaws.com" {
						continue
					}
					if *e.EventName == "CreateStack" {
						stack.Creator = *e.Username
					}
				}

				if _, ok := stacksByName[*tag.Value]; !ok {
					stacksByName[*tag.Value] = []stackDescriptor{}
				}
				stacksByName[*tag.Value] = append(stacksByName[*tag.Value], stack)
			}
		}
	}

	body, err := json.Marshal(stacksByName)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 502}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(doListCluster)
}
