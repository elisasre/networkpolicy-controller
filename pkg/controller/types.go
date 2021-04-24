package controller

import (
	networkv1 "k8s.io/api/networking/v1"
)

type networkRule struct {
	IgnoredNamespaces []string                 `yaml:"ignoredNamespaces"`
	Spec              *networkv1.NetworkPolicy `yaml:"spec"`
}

// Config ...
type Config struct {
	rules            []networkRule `yaml:"rules"`
	ignoreAnnotation string        `yaml:"ignoreAnnotation"`
}
