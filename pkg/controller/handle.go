package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func notFound(err error, name string) bool {
	match, err := regexp.MatchString(fmt.Sprintf("\"%s\" not found$", name), err.Error())
	if err != nil {
		return false
	}
	return match
}

func (c *Controller) ensurePoliciesExist(ns *v1.Namespace) {
	for _, rule := range c.config.Rules {
		if Contains(rule.IgnoredNamespaces, ns.Name) {
			continue
		}
		val, ok := ns.ObjectMeta.Annotations[c.config.IgnoreAnnotation]
		if ok && val == "true" {
			continue
		}
		asBytes, err := json.Marshal(rule.Spec)
		if err != nil {
			log.Printf("Got error while marshalling %s: %v", ns.Name, err)
			continue
		}
		spec := strings.ReplaceAll(string(asBytes), "self()", ns.Name)
		finalSpec := &networkv1.NetworkPolicy{}
		err = json.Unmarshal([]byte(spec), &finalSpec)
		if err != nil {
			log.Printf("Got error while unmarshalling %s: %v", ns.Name, err)
			continue
		}
		ctx := context.Background()
		_, err = c.kclient.NetworkingV1().NetworkPolicies(ns.Name).Get(ctx, finalSpec.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				_, errCreate := c.kclient.NetworkingV1().NetworkPolicies(ns.Name).Create(ctx, finalSpec, metav1.CreateOptions{})
				if errCreate != nil {
					log.Printf("Got error while creating networkpolicy %s - %s: %v", ns.Name, finalSpec.Name, errCreate)
				}
			} else {
				log.Printf("Got error while finding networkpolicy %s: %v", ns.Name, err)
			}
		}
	}
}
