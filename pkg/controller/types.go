package controller

import (
	networkv1 "k8s.io/api/networking/v1"
)

// NetworkRule ...
type NetworkRule struct {
	IgnoredNamespaces []string                 `json:"ignoredNamespaces"`
	Spec              *networkv1.NetworkPolicy `json:"spec"`
}

// Config ...
type Config struct {
	Rules            []NetworkRule `json:"rules"`
	IgnoreAnnotation string        `json:"ignoreAnnotation"`
}
