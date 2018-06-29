package seaotter

import (
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
// Controller implementation logic for SeaOtter resources goes here.

func (bc *SeaOtterController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on seaotter.SeaOtterController to reconcile %s\n", k.Name)
	return nil
}

// +kubebuilder:controller:group=otters,version=v1alpha1,kind=SeaOtter,resource=seaotters
type SeaOtterController struct {
	// INSERT ADDITIONAL FIELDS HERE
	seaotterLister ottersv1alpha1lister.SeaOtterLister
	seaotterclient ottersv1alpha1client.OttersV1alpha1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	seaotterrecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &SeaOtterController{
		seaotterLister: arguments.ControllerManager.GetInformerProvider(&ottersv1alpha1.SeaOtter{}).(ottersv1alpha1informer.SeaOtterInformer).Lister(),

		seaotterclient:   arguments.Clientset.OttersV1alpha1(),
		seaotterrecorder: arguments.CreateRecorder("SeaOtterController"),
	}

	// Create a new controller that will call SeaOtterController.Reconcile on changes to SeaOtters
	gc := &controller.GenericController{
		Name:             "SeaOtterController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&ottersv1alpha1.SeaOtter{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a SeaOtter Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the SeaOtterController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
