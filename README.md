# networkpolicy-controller

This controller makes sure that following policy exist in all namespaces except `kube-system`:

```
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-all-except-metadata
spec:
  podSelector: {}
  egress:
  - to:
    - ipBlock:
        cidr: 0.0.0.0/0
        except: 
        - 169.254.169.254/32
  policyTypes:
  - Egress
```

In other words, this component blocks access to common metadata ip address in different cloudproviders.
