# networkpolicy-controller

This component adds common networkpolicies in namespace levels.
All rules can be configured using config file.

**Example config**

```
{
  "ignoreAnnotation": "networkpolicy.elisa.fi",
  "rules": [
    {
      "ignoredNamespaces": [
        "kube-system",
        "ingress-nginx-internal",
        "ingress-nginx-external",
        "monitoring"
      ],
      "spec": {
        "kind": "NetworkPolicy",
        "apiVersion": "networking.k8s.io/v1",
        "metadata": {
          "name": "namespace-isolation"
        },
        "spec": {
          "podSelector": {},
          "ingress": [
            {
              "from": [
                {
                  "namespaceSelector": {
                    "matchLabels": {
                      "app.kubernetes.io/name": "self()"
                    }
                  }
                },
                {
                  "namespaceSelector": {
                    "matchLabels": {
                      "app.kubernetes.io/name": "ingress-nginx-internal"
                    }
                  }
                },
                {
                  "namespaceSelector": {
                    "matchLabels": {
                      "app.kubernetes.io/name": "ingress-nginx-external"
                    }
                  }
                },
                {
                  "namespaceSelector": {
                    "matchLabels": {
                      "app.kubernetes.io/name": "monitoring"
                    }
                  }
                }
              ]
            }
          ]
        }
      }
    }
  ]
}
```

This config means that it will add following networkpolicy to all namespaces
(excluding `kube-system`, `ingress-nginx-internal`, `ingress-nginx-external`,
`monitoring` and namespaces which do have annotation `networkpolicy.elisa.fi` and value `true`).
`self()` will be replaced with the current namespace name.

```
kind: NetworkPolicy
apiVersion: networking.k8s.io/v1
metadata:
  name: namespace-isolation
spec:
  podSelector: {}
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          app.kubernetes.io/name: self()
    - namespaceSelector:
        matchLabels:
          app.kubernetes.io/name: ingress-nginx-internal
    - namespaceSelector:
        matchLabels:
          app.kubernetes.io/name: ingress-nginx-external
    - namespaceSelector:
        matchLabels:
          app.kubernetes.io/name: monitoring
```

# Install

`kubectl apply -f https://raw.githubusercontent.com/elisasre/networkpolicy-controller/master/manifests/deploy.yaml`
