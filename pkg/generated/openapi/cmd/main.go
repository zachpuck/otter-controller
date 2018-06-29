package main

import (
	"github.com/kubernetes-sigs/kubebuilder/pkg/docs"
	"github.com/zachpuck/otter-controller/pkg/generated/openapi"
)

func main() {
	docs.WriteOpenAPI(openapi.GetOpenAPIDefinitions)
}
