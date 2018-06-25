// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/zachpuck/otter-controller/pkg/apis/otters
// +k8s:defaulter-gen=TypeMeta
// +groupName=otters.k8s.dokuforest.com
package v1alpha1 // import "github.com/zachpuck/otter-controller/pkg/apis/otters/v1alpha1"
