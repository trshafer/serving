/*
Copyright 2020 The Knative Authors
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

package ingress

import (
	"knative.dev/serving/pkg/apis/serving/v1alpha1"
)

const (
	// IstioCanonicalServiceLabelName is the name of label for the Istio Canonical Service for a workload instance.
	IstioCanonicalServiceLabelName = "service.istio.io/canonical-name"

	// IstioCanonicalServiceRevisionLabelName is the name of label for the Istio Canonical Service revision for a workload instance.
	IstioCanonicalServiceRevisionLabelName = "service.istio.io/canonical-revision"

	// ReplaceWithRevisionName is the a placeholder for being replaced by the configuration.
	ReplaceWithRevisionName = "knative.dev.placeholder-revision-name"
)

// AddLabels adds labels
func AddLabels(service *v1alpha1.Service) *v1alpha1.Service {
	// newService := service.DeepCopy()

	serviceName := service.Name
	// if service.ObjectMeta.Labels == nil {
	// 	service.ObjectMeta.Labels = make(map[string]string)
	// }
	// // service.ObjectMeta.Labels
	// service.ObjectMeta.Labels[IstioCanonicalServiceLabelName] = serviceName
	// service.ObjectMeta.Labels[IstioCanonicalServiceRevisionLabelName] = ReplaceWithRevisionName

	if service.Spec.Template.ObjectMeta.Labels == nil {
		service.Spec.Template.ObjectMeta.Labels = make(map[string]string)
	}

	service.Spec.Template.ObjectMeta.Labels[IstioCanonicalServiceLabelName] = serviceName
	service.Spec.Template.ObjectMeta.Labels[IstioCanonicalServiceRevisionLabelName] = ReplaceWithRevisionName
	// newService.SetDefaults
	// newService.Spec.Template.GetLabels()[IstioCanonicalServiceLabelName] = serviceName
	// newService.Spec.Template.GetLabels()[IstioCanonicalServiceRevisionLabelName] = ReplaceWithRevisionName
	// newService.Spec.Template.Labels[IstioCanonicalServiceLabelName] = serviceName
	// newService.Spec.Template.Labels[IstioCanonicalServiceRevisionLabelName] = ReplaceWithRevisionName

	return service
}
