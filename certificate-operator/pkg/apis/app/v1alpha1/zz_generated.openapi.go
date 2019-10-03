// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/app/v1alpha1.IstioCertificate":       schema_pkg_apis_app_v1alpha1_IstioCertificate(ref),
		"./pkg/apis/app/v1alpha1.IstioCertificateSpec":   schema_pkg_apis_app_v1alpha1_IstioCertificateSpec(ref),
		"./pkg/apis/app/v1alpha1.IstioCertificateStatus": schema_pkg_apis_app_v1alpha1_IstioCertificateStatus(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_IstioCertificate(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IstioCertificate is the Schema for the istiocertificates API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/app/v1alpha1.IstioCertificateSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/app/v1alpha1.IstioCertificateStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/app/v1alpha1.IstioCertificateSpec", "./pkg/apis/app/v1alpha1.IstioCertificateStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_IstioCertificateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IstioCertificateSpec defines the desired state of IstioCertificate",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Unique name of resource",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"hosts": {
						SchemaProps: spec.SchemaProps{
							Description: "List of Servers > map of list of hosts and port",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"mode": {
						SchemaProps: spec.SchemaProps{
							Description: "Options: SIMPLE|PASSTHROUGH|MUTUAL",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"protocol": {
						SchemaProps: spec.SchemaProps{
							Description: "Options: HTTP|HTTPS|GRPC|HTTP2|MONGO|TCP",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"trafficType": {
						SchemaProps: spec.SchemaProps{
							Description: "Options: \"ingress\" or \"egress\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tlsSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies TLS Cert/Key to be created",
							Ref:         ref("./pkg/apis/app/v1alpha1.TLSSecret"),
						},
					},
					"tlsSecretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the TLS Secret",
							Ref:         ref("./pkg/apis/app/v1alpha1.TLSSecretRef"),
						},
					},
					"tlsSecretPath": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies TLS Cert/Key Path if not using SDS",
							Ref:         ref("./pkg/apis/app/v1alpha1.TLSSecretPath"),
						},
					},
				},
				Required: []string{"name", "hosts", "port", "mode", "protocol", "trafficType", "tlsSecret", "tlsSecretRef", "tlsSecretPath"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/app/v1alpha1.TLSSecret", "./pkg/apis/app/v1alpha1.TLSSecretPath", "./pkg/apis/app/v1alpha1.TLSSecretRef"},
	}
}

func schema_pkg_apis_app_v1alpha1_IstioCertificateStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "IstioCertificateStatus defines the observed state of IstioCertificate",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
