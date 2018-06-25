package otter

import (
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/eventhandlers"
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/predicates"
	"k8s.io/client-go/tools/record"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

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

	// Read the Otter state
	ot, err := bc.Clientset.
		OttersV1alpha1().
		Otters(k.Namespace).
		Get(k.Name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
	}

	// Create the cononical DeploymentSpec
	spec := appsv1.DeploymentSpec {
		Selector: &metav1.LabelSelector {
			MatchLabels: map[string]string {
				"otter": k.Name},
		},
		Replicas: &ot.Spec.Replicas,
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"otter": k.Name},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name: k.Name,
						Image: ot.Spec.Image,
					},
				},
			},
		},
	}

	// Read the DeploymentState
	dep, err := bc.KubernetesClientSet.
		AppsV1().
		Deployments(k.Namespace).
		Get(k.Name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		// Create the Deployment
		dep = &appsv1.Deployment{
			Spec: spec,
		}
		// Set OwnerReferences so the Deployment is GCed
		dep.OwnerReferences = []metav1.OwnerReference{
			*metav1.NewControllerRef(ot, schema.GroupVersionKind{
				Group: "k8s.dokuforest.com",
				Version: "v1alpha1",
				Kind: "Otter",
			}),
		}
		dep.Name = k.Name
		dep.Namespace = k.Namespace
		_, err = bc.KubernetesClientSet.AppsV1().
			Deployments(k.Namespace).Create(dep)
	} else {
		// Update the Deployment if its observed Spec does not match the desired Spec
		image := dep.Spec.Template.Spec.Containers[0].Image
		replicas := *dep.Spec.Replicas
		if replicas == ot.Spec.Replicas && image == ot.Spec.Image {
			return nil
		}
		dep.Name = k.Name
		dep.Namespace = k.Namespace
		dep.Spec = spec
		_, err = bc.KubernetesClientSet.AppsV1().
			Deployments(k.Namespace).Update(dep)
	}
	if err != nil {
		return err
	}

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
			OttersV1alpha1().
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
