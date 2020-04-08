package parsetmpl

import (
	"strings"
)

type Info struct {
	Title       string `yaml:"title"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}
type Feature struct {
	FeatureName string              `yaml:"feature"`
	ID          string              `yaml:"id"`
	Description string              `yaml:"description"`
	Endpoints   []map[string]string `yaml:"endpoints"`
}
type Method struct {
	Description string                 `yaml:"description"`
	OperationID string                 `yaml:"operationId"`
	Responses   map[string]interface{} `yaml:"responses"`
	RBACFeature []string               `yaml:"x-rbac-feature"`
}
type Server struct {
	URL string "yaml:`url`"
}
type OpenApiSpec struct {
	Info    Info                         `yaml:"info"`
	Servers []Server                     `yaml:"servers"`
	Paths   map[string]map[string]Method `yaml:"paths"`
}
type Result struct {
	ServiceName string    `yaml:"serviceName"`
	Features    []Feature `yaml:"features"`
}

func (openApiSpec *OpenApiSpec) Parsetmpl() Result {
	var result Result
	// Parse file
	result.ServiceName = openApiSpec.Info.Title

	for pathKey, path := range openApiSpec.Paths {
		for methodKey, method := range path {
			for _, feature := range method.RBACFeature {

				var endpoints []map[string]string
				endpoint := make(map[string]string)
				endpoint[pathKey] = methodKey
				endpoints = append(endpoints, endpoint)

				result.Features = append(result.Features, Feature{
					FeatureName: strings.ReplaceAll(feature, "_", " "),
					ID:          feature,
					Description: method.Description,
					Endpoints:   endpoints,
				})
			}
		}
	}
	return result
}
