package resources

import (
	"fmt"
	"net"
	"strings"

	"github.com/jenkinsci/kubernetes-operator/api/v1alpha2"
	"github.com/jenkinsci/kubernetes-operator/pkg/constants"
	stackerr "github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
)

// ServiceKind the kind name for Service
const ServiceKind = "Service"

// UpdateService returns new service with override fields from config
func UpdateService(actual corev1.Service, config v1alpha2.Service) corev1.Service {
	actual.ObjectMeta.Annotations = config.Annotations
	for key, value := range config.Labels {
		actual.ObjectMeta.Labels[key] = value
	}
	actual.Spec.Type = config.Type
	actual.Spec.LoadBalancerIP = config.LoadBalancerIP
	actual.Spec.LoadBalancerSourceRanges = config.LoadBalancerSourceRanges
	if len(actual.Spec.Ports) == 0 {
		actual.Spec.Ports = []corev1.ServicePort{{}}
	}
	actual.Spec.Ports[0].Port = config.Port
	if config.NodePort != 0 {
		actual.Spec.Ports[0].NodePort = config.NodePort
	}

	return actual
}

// GetJenkinsHTTPServiceName returns Kubernetes service name used for expose Jenkins HTTP endpoint
func GetJenkinsHTTPServiceName(jenkins *v1alpha2.Jenkins) string {
	return fmt.Sprintf("%s-%s", constants.LabelAppValue, jenkins.ObjectMeta.Name)
}

// GetJenkinsJNLPServiceName returns Kubernetes service name used for expose Jenkins JNLP endpoint
func GetJenkinsJNLPServiceName(jenkins *v1alpha2.Jenkins) string {
	return fmt.Sprintf("%s-%s-jnlp", constants.LabelAppValue, jenkins.ObjectMeta.Name)
}

// GetJenkinsHTTPServiceFQDN returns Kubernetes service FQDN used for expose Jenkins HTTP endpoint
func GetJenkinsHTTPServiceFQDN(jenkins *v1alpha2.Jenkins) (string, error) {
	clusterDomain, err := getClusterDomain()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s.%s.svc.%s", constants.LabelAppValue, jenkins.ObjectMeta.Name, jenkins.ObjectMeta.Namespace, clusterDomain), nil
}

// GetJenkinsJNLPServiceFQDN returns Kubernetes service FQDN used for expose Jenkins JNLP endpoint
func GetJenkinsJNLPServiceFQDN(jenkins *v1alpha2.Jenkins) (string, error) {
	clusterDomain, err := getClusterDomain()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s-jnlp.%s.svc.%s", constants.LabelAppValue, jenkins.ObjectMeta.Name, jenkins.ObjectMeta.Namespace, clusterDomain), nil
}

// GetClusterDomain returns Kubernetes cluster domain, default to "cluster.local"
func getClusterDomain() (string, error) {
	apiSvc := "kubernetes.default.svc"
	cname, err := net.LookupCNAME(apiSvc)
	if err != nil {
		return "", stackerr.WithStack(err)
	}

	clusterDomain := strings.TrimPrefix(cname, "kubernetes.default.svc")
	clusterDomain = strings.TrimPrefix(clusterDomain, ".")
	clusterDomain = strings.TrimSuffix(clusterDomain, ".")

	return clusterDomain, nil
}
