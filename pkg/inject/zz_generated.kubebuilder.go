package inject

import (
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	ottersv1alpha1 "github.com/zachpuck/otter-controller/pkg/apis/otters/v1alpha1"
	rscheme "github.com/zachpuck/otter-controller/pkg/client/clientset/versioned/scheme"
	"github.com/zachpuck/otter-controller/pkg/controller/otter"
	"github.com/zachpuck/otter-controller/pkg/inject/args"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&ottersv1alpha1.Otter{}, arguments.Informers.Otters().V1alpha1().Otters()); err != nil {
			return err
		}

		// Add Kubernetes informers

		if c, err := otter.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &ottersv1alpha1.OtterCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"otters.k8s.dokuforest.com"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "otters.k8s.dokuforest.com",
		Version: "v1alpha1",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
