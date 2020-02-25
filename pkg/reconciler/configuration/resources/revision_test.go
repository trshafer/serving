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

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/ptr"
	"knative.dev/serving/pkg/apis/networking"
	"knative.dev/serving/pkg/apis/serving"
	v1 "knative.dev/serving/pkg/apis/serving/v1"
)

func TestMakeRevisions(t *testing.T) {
	tests := []struct {
		name          string
		configuration *v1.Configuration
		want          *v1.Revision
	}{{
		name: "no build",
		configuration: &v1.Configuration{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:  "no",
				Name:       "build",
				Generation: 10,
			},
			Spec: v1.ConfigurationSpec{
				Template: v1.RevisionTemplateSpec{
					Spec: v1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Image: "busybox",
							}},
						},
					},
				},
			},
		},
		want: &v1.Revision{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    "no",
				GenerateName: "build-",
				Annotations:  map[string]string{},
				OwnerReferences: []metav1.OwnerReference{{
					APIVersion:         v1.SchemeGroupVersion.String(),
					Kind:               "Configuration",
					Name:               "build",
					Controller:         ptr.Bool(true),
					BlockOwnerDeletion: ptr.Bool(true),
				}},
				Labels: map[string]string{
					serving.ConfigurationLabelKey:           "build",
					serving.ConfigurationGenerationLabelKey: "10",
					serving.ServiceLabelKey:                 "",
					networking.IngressClassLabelKey:         "",
				},
			},
			Spec: v1.RevisionSpec{
				PodSpec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "busybox",
					}},
				},
			},
		},
	}, {
		name: "with labels",
		configuration: &v1.Configuration{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:  "with",
				Name:       "labels",
				Generation: 100,
			},
			Spec: v1.ConfigurationSpec{
				Template: v1.RevisionTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"foo": "bar",
							"baz": "blah",
						},
					},
					Spec: v1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Image: "busybox",
							}},
						},
					},
				},
			},
		},
		want: &v1.Revision{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    "with",
				GenerateName: "labels-",
				Annotations:  map[string]string{},
				OwnerReferences: []metav1.OwnerReference{{
					APIVersion:         v1.SchemeGroupVersion.String(),
					Kind:               "Configuration",
					Name:               "labels",
					Controller:         ptr.Bool(true),
					BlockOwnerDeletion: ptr.Bool(true),
				}},
				Labels: map[string]string{
					serving.ConfigurationLabelKey:           "labels",
					serving.ConfigurationGenerationLabelKey: "100",
					serving.ServiceLabelKey:                 "",
					networking.IngressClassLabelKey:         "",
					"foo":                                   "bar",
					"baz":                                   "blah",
				},
			},
			Spec: v1.RevisionSpec{
				PodSpec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "busybox",
					}},
				},
			},
		},
	}, {
		name: "with networking label",
		configuration: &v1.Configuration{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:  "with",
				Name:       "labels",
				Generation: 100,
				Labels: map[string]string{
					networking.IngressClassLabelKey: "test-ingress-class",
				},
			},
			Spec: v1.ConfigurationSpec{
				Template: v1.RevisionTemplateSpec{
					Spec: v1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Image: "busybox",
							}},
						},
					},
				},
			},
		},
		want: &v1.Revision{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    "with",
				GenerateName: "labels-",
				Annotations:  map[string]string{},
				OwnerReferences: []metav1.OwnerReference{{
					APIVersion:         v1.SchemeGroupVersion.String(),
					Kind:               "Configuration",
					Name:               "labels",
					Controller:         ptr.Bool(true),
					BlockOwnerDeletion: ptr.Bool(true),
				}},
				Labels: map[string]string{
					serving.ConfigurationLabelKey:           "labels",
					serving.ConfigurationGenerationLabelKey: "100",
					serving.ServiceLabelKey:                 "",
					networking.IngressClassLabelKey:         "test-ingress-class",
				},
			},
			Spec: v1.RevisionSpec{
				PodSpec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "busybox",
					}},
				},
			},
		},
	}, {
		name: "with annotations",
		configuration: &v1.Configuration{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:  "with",
				Name:       "annotations",
				Generation: 100,
			},
			Spec: v1.ConfigurationSpec{
				Template: v1.RevisionTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"foo": "bar",
							"baz": "blah",
						},
					},
					Spec: v1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Image: "busybox",
							}},
						},
					},
				},
			},
		},
		want: &v1.Revision{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    "with",
				GenerateName: "annotations-",
				OwnerReferences: []metav1.OwnerReference{{
					APIVersion:         v1.SchemeGroupVersion.String(),
					Kind:               "Configuration",
					Name:               "annotations",
					Controller:         ptr.Bool(true),
					BlockOwnerDeletion: ptr.Bool(true),
				}},
				Labels: map[string]string{
					serving.ConfigurationLabelKey:           "annotations",
					serving.ConfigurationGenerationLabelKey: "100",
					serving.ServiceLabelKey:                 "",
					networking.IngressClassLabelKey:         "",
				},
				Annotations: map[string]string{
					"foo": "bar",
					"baz": "blah",
				},
			},
			Spec: v1.RevisionSpec{
				PodSpec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "busybox",
					}},
				},
			},
		},
	}, {
		name: "with creator annotation from config",
		configuration: &v1.Configuration{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "anno",
				Name:      "config",
				Annotations: map[string]string{
					"serving.knative.dev/creator":      "admin",
					"serving.knative.dev/lastModifier": "someone",
				},
				Generation: 10,
			},
			Spec: v1.ConfigurationSpec{
				Template: v1.RevisionTemplateSpec{
					Spec: v1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Image: "busybox",
							}},
						},
					},
				},
			},
		},
		want: &v1.Revision{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    "anno",
				GenerateName: "config-",
				Annotations: map[string]string{
					"serving.knative.dev/creator": "someone",
				},
				OwnerReferences: []metav1.OwnerReference{{
					APIVersion:         v1.SchemeGroupVersion.String(),
					Kind:               "Configuration",
					Name:               "config",
					Controller:         ptr.Bool(true),
					BlockOwnerDeletion: ptr.Bool(true),
				}},
				Labels: map[string]string{
					serving.ConfigurationLabelKey:           "config",
					serving.ConfigurationGenerationLabelKey: "10",
					serving.ServiceLabelKey:                 "",
					networking.IngressClassLabelKey:         "",
				},
			},
			Spec: v1.RevisionSpec{
				PodSpec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "busybox",
					}},
				},
			},
		},
	}, {
		name: "with creator annotation from config with other annotations",
		configuration: &v1.Configuration{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "anno",
				Name:      "config",
				Annotations: map[string]string{
					"serving.knative.dev/creator":      "admin",
					"serving.knative.dev/lastModifier": "someone",
				},
				Generation: 10,
			},
			Spec: v1.ConfigurationSpec{
				Template: v1.RevisionTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Annotations: map[string]string{
							"foo": "bar",
							"baz": "blah",
						},
					},
					Spec: v1.RevisionSpec{
						PodSpec: corev1.PodSpec{
							Containers: []corev1.Container{{
								Image: "busybox",
							}},
						},
					},
				},
			},
		},
		want: &v1.Revision{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    "anno",
				GenerateName: "config-",
				Annotations: map[string]string{
					"serving.knative.dev/creator": "someone",
					"foo":                         "bar",
					"baz":                         "blah",
				},
				OwnerReferences: []metav1.OwnerReference{{
					APIVersion:         v1.SchemeGroupVersion.String(),
					Kind:               "Configuration",
					Name:               "config",
					Controller:         ptr.Bool(true),
					BlockOwnerDeletion: ptr.Bool(true),
				}},
				Labels: map[string]string{
					serving.ConfigurationLabelKey:           "config",
					serving.ConfigurationGenerationLabelKey: "10",
					serving.ServiceLabelKey:                 "",
					networking.IngressClassLabelKey:         "",
				},
			},
			Spec: v1.RevisionSpec{
				PodSpec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "busybox",
					}},
				},
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MakeRevision(test.configuration)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("MakeRevision (-want, +got) = %v", diff)
			}
		})
	}
}
