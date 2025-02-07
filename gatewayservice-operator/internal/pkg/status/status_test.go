package status_test

import (
	"fmt"
	"reflect"
	"testing"

	s "github.com/xunholy/k8s-istio-gateway-service-operator/internal/pkg/status"
	appv1alpha1 "github.com/xunholy/k8s-istio-gateway-service-operator/pkg/apis/crd/v1alpha1"
)

var (
	name      = "example-app"
	namespace = "application"
)

func TestStatusReconcile(t *testing.T) {
	expected := &appv1alpha1.GatewayServiceStatus{
		Condition: appv1alpha1.Condition{
			Success:      true,
			ErrorMessage: "",
			CreatedSecretDetails: appv1alpha1.CreatedSecretDetails{
				SecretName:      fmt.Sprintf("%s-%s-secret", name, namespace),
				SecretNamespace: namespace,
			},
		},
	}
	statusConfig := s.StatusConfig{
		Success:         true,
		ErrorMessage:    "",
		SecretName:      fmt.Sprintf("%s-%s-secret", name, namespace),
		SecretNamespace: namespace,
	}
	secretObject := s.Reconcile(statusConfig)
	if !reflect.DeepEqual(secretObject, expected) {
		t.Fatalf("Expected: (%+v)\n Found: (%+v)", expected, secretObject)
	}
}
