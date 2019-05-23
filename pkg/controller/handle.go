package controller

import (
	"fmt"
	"log"
	"regexp"

	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kubeSystem           = "kube-system"
	cloudMetadataAddress = "169.254.169.254/32"
)

var policy = &networkv1.NetworkPolicy{
	ObjectMeta: metav1.ObjectMeta{
		Name: "allow-all-except-metadata",
	},
	Spec: networkv1.NetworkPolicySpec{
		PolicyTypes: []networkv1.PolicyType{networkv1.PolicyTypeEgress},
		Egress: []networkv1.NetworkPolicyEgressRule{
			{
				To: []networkv1.NetworkPolicyPeer{
					{
						IPBlock: &networkv1.IPBlock{
							CIDR:   "0.0.0.0/0",
							Except: []string{cloudMetadataAddress},
						},
					},
				},
			},
		},
	},
}

func notFound(err error, name string) bool {
	match, err := regexp.MatchString(fmt.Sprintf("\"%s\" not found$", name), err.Error())
	if err != nil {
		return false
	}
	return match
}

func (c *Controller) ensurePolicyExist(namespace string) {
	if namespace != kubeSystem {
		_, err := c.kclient.NetworkingV1().NetworkPolicies(namespace).Update(policy)
		if err != nil {
			if notFound(err, policy.Name) {
				_, err = c.kclient.NetworkingV1().NetworkPolicies(namespace).Create(policy)
			}
			if err != nil {
				log.Printf("ERROR namespace %s: %v", namespace, err)
			}
		}
	}
}
