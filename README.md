# networkpolicy-controller

This component adds common networkpolicies in namespace levels.
All rules can be configured using config file.

**Example config**

```
ignoreAnnotation: networkpolicy.elisa.fi
rules:
- ignoredNamespaces:
  - kube-system
  spec:
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

This config means that it will add following networkpolicy to all namespaces
(excluding `kube-system` and namespaces which do have annotation `networkpolicy.elisa.fi` and value `true`)

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
