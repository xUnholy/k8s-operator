package secret_test

import (
	"fmt"
	"reflect"
	"testing"

	s "github.com/xUnholy/k8s-operator/internal/pkg/secret"
	appv1alpha1 "github.com/xUnholy/k8s-operator/pkg/apis/app/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	name      = "example-app"
	namespace = "application"
)

func TestSecretReconcile(t *testing.T) {
	certificate := &appv1alpha1.IstioCertificate{
		Spec: appv1alpha1.IstioCertificateSpec{
			TLSOptions: appv1alpha1.TLSOptions{
				TLSSecret: &appv1alpha1.TLSSecret{
					Cert: []byte{1, 2},
					Key:  []byte{1, 2},
				},
			},
		},
	}
	expected := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%s-secret", name, namespace),
			Namespace: namespace,
			Labels:    map[string]string{"Namespace": namespace},
		},
		Data: map[string][]byte{
			"tls.key": []byte{1, 2},
			"tls.crt": []byte{1, 2},
		},
		Type: "kubernetes.io/tls",
	}
	secretConfig := s.SecretConfig{
		Name:        fmt.Sprintf("%s-%s-secret", name, namespace),
		Namespace:   namespace,
		Labels:      map[string]string{"Namespace": namespace},
		Certificate: certificate,
	}
	secretObject := s.Reconcile(secretConfig)
	if !reflect.DeepEqual(secretObject, expected) {
		t.Fatalf("Expected: (%+v)\n Found: (%+v)", expected, secretObject)
	}
}
