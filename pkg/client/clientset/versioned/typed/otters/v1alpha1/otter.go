// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zachpuck/otter-controller/pkg/apis/otters/v1alpha1"
	scheme "github.com/zachpuck/otter-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OttersGetter has a method to return a OtterInterface.
// A group's client should implement this interface.
type OttersGetter interface {
	Otters(namespace string) OtterInterface
}

// OtterInterface has methods to work with Otter resources.
type OtterInterface interface {
	Create(*v1alpha1.Otter) (*v1alpha1.Otter, error)
	Update(*v1alpha1.Otter) (*v1alpha1.Otter, error)
	UpdateStatus(*v1alpha1.Otter) (*v1alpha1.Otter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Otter, error)
	List(opts v1.ListOptions) (*v1alpha1.OtterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Otter, err error)
	OtterExpansion
}

// otters implements OtterInterface
type otters struct {
	client rest.Interface
	ns     string
}

// newOtters returns a Otters
func newOtters(c *OttersV1alpha1Client, namespace string) *otters {
	return &otters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the otter, and returns the corresponding otter object, and an error if there is any.
func (c *otters) Get(name string, options v1.GetOptions) (result *v1alpha1.Otter, err error) {
	result = &v1alpha1.Otter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("otters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Otters that match those selectors.
func (c *otters) List(opts v1.ListOptions) (result *v1alpha1.OtterList, err error) {
	result = &v1alpha1.OtterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("otters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested otters.
func (c *otters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("otters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a otter and creates it.  Returns the server's representation of the otter, and an error, if there is any.
func (c *otters) Create(otter *v1alpha1.Otter) (result *v1alpha1.Otter, err error) {
	result = &v1alpha1.Otter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("otters").
		Body(otter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a otter and updates it. Returns the server's representation of the otter, and an error, if there is any.
func (c *otters) Update(otter *v1alpha1.Otter) (result *v1alpha1.Otter, err error) {
	result = &v1alpha1.Otter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("otters").
		Name(otter.Name).
		Body(otter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *otters) UpdateStatus(otter *v1alpha1.Otter) (result *v1alpha1.Otter, err error) {
	result = &v1alpha1.Otter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("otters").
		Name(otter.Name).
		SubResource("status").
		Body(otter).
		Do().
		Into(result)
	return
}

// Delete takes name of the otter and deletes it. Returns an error if one occurs.
func (c *otters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("otters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *otters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("otters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched otter.
func (c *otters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Otter, err error) {
	result = &v1alpha1.Otter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("otters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}