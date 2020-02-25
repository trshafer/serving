/*
Copyright 2018 The Knative Authors

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

package resources

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"knative.dev/serving/pkg/apis/networking"
	"knative.dev/serving/pkg/apis/serving"
)

func TestConfigurationSpec(t *testing.T) {
	s := createService()
	c, _ := MakeConfiguration(s, "test-ingress-class")
	if got, want := c.Name, testServiceName; got != want {
		t.Errorf("expected %q for service name got %q", want, got)
	}
	if got, want := c.Namespace, testServiceNamespace; got != want {
		t.Errorf("expected %q for service namespace got %q", want, got)
	}
	if got, want := c.Spec.GetTemplate().Spec.GetContainer().Name, testContainerName; got != want {
		t.Errorf("expected %q for container name got %q", want, got)
	}
	expectOwnerReferencesSetCorrectly(t, c.OwnerReferences)

	if got, want := len(c.Labels), 3; got != want {
		t.Errorf("expected %d labels got %d", want, got)
	}
	if got, want := c.Labels[serving.ServiceLabelKey], testServiceName; got != want {
		t.Errorf("expected %q labels got %q", want, got)
	}
	if got, want := c.Labels[networking.IngressClassLabelKey], "test-ingress-class"; got != want {
		t.Errorf("expected %q labels got %q", want, got)
	}
}

func TestConfigurationHasNoKubectlAnnotation(t *testing.T) {
	s := createServiceWithKubectlAnnotation()
	c, err := MakeConfiguration(s, "test-ingress-class")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if v, ok := c.Annotations[corev1.LastAppliedConfigAnnotation]; ok {
		t.Errorf("Annotation %s = %q, want empty", corev1.LastAppliedConfigAnnotation, v)
	}
}
