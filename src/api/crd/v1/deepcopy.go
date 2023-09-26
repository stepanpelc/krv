/*
    krv - kubernetes resource validator
    Copyright (C) 2022 SIZEK s.r.o

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
