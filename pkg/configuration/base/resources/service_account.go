package resources

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewServiceAccount return Kubernetes service account
func NewServiceAccount(meta metav1.ObjectMeta, annotations map[string]string) *v1.ServiceAccount {
	meta.Annotations = annotations
	return &v1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: meta,
	}
}
