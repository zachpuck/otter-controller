package otter

import (
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/eventhandlers"
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	ottersv1alpha1 "github.com/zachpuck/otter-controller/pkg/apis/otters/v1alpha1"
	ottersv1alpha1client "github.com/zachpuck/otter-controller/pkg/client/clientset/versioned/typed/otters/v1alpha1"
	ottersv1alpha1informer "github.com/zachpuck/otter-controller/pkg/client/informers/externalversions/otters/v1alpha1"
	ottersv1alpha1lister "github.com/zachpuck/otter-controller/pkg/client/listers/otters/v1alpha1"

	"github.com/zachpuck/otter-controller/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for Otter resources goes here.

func (bc *OtterController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on otter.OtterController to reconcile %s\n", k.Name)
	return nil
}

// OtterController is a sample of implementing a custom controller
// +kubebuilder:controller:group=otters,version=v1alpha1,kind=Otter,resource=otters
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:informers:group=apps,version=v1,kind=Deployment
type OtterController struct {
	// InjectArgs contains the clients provided to ProvideController
	args.InjectArgs

	otterLister ottersv1alpha1lister.OtterLister
	otterclient ottersv1alpha1client.OttersV1alpha1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	otterrecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &OtterController{
		InjectArgs: arguments,
		otterLister: arguments.ControllerManager.GetInformerProvider(&ottersv1alpha1.Otter{}).(ottersv1alpha1informer.OtterInformer).Lister(),

		otterclient:   arguments.Clientset.OttersV1alpha1(),
		otterrecorder: arguments.CreateRecorder("OtterController"),
	}

	// Create a new controller that will call OtterController.Reconcile on changes to Otters
	gc := &controller.GenericController{
		Name:             "OtterController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	// Watch Otter
	if err := gc.Watch(&ottersv1alpha1.Otter{}); err != nil {
		return gc, err
	}

	// Watch Deployments
	otterLookup := func(k types.ReconcileKey) (interface{}, error) {
		d, err := bc.Clientset.
			WorkloadsV1alpha1().
			Otters(k.Namespace).
			Get(k.Name, metav1.GetOptions{})
		return d, err
	}
	if err := gc.WatchControllerOf(
		&appsv1.Deployment{},
		eventhandlers.Path{otterLookup},
		predicates.ResourceVersionChanged); err != nil {
			return gc, err
	}
	
	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a Otter Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the OtterController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
