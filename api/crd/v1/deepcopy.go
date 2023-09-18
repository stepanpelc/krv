package v1

//deep copy for our CRD

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *Validation) DeepCopyInto(out *Validation) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = ValidationSpec{
		Name:      in.Spec.Name,
		Namespace: in.Spec.Namespace,
		Resource:  in.Spec.Resource,
	}
	out.Status = StatusInfo{
		LastCheck:   in.Status.LastCheck,
		LastChanged: in.Status.LastChanged,
		State:       in.Status.State,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *Validation) DeepCopyObject() runtime.Object {
	out := Validation{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *ValidationList) DeepCopyObject() runtime.Object {
	out := ValidationList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Validation, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
