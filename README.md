# networkpolicy-controller

This component blocks access to common metadata IP address used within cloud providers by forcing the following network policy into all namespaces except `kube-system`.

**Example policy**

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

# Install

`kubectl apply -f https://raw.githubusercontent.com/ElisaOyj/networkpolicy-controller/master/manifests/deploy.yaml`
