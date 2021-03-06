/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubernetes

// This is copied from k8s.io/kubernetes to avoid a dependency on all of Kubernetes.
// TODO(enisoc): Move the upstream code to somewhere better.

// BoolPtr returns a pointer to a bool
func BoolPtr(b bool) *bool {
	o := b
	return &o
}

// Int32Ptr returns a pointer to an int32
func Int32Ptr(i int32) *int32 {
	o := i
	return &o
}
