# AppOperator
A demo for validation of specific crd

## Requirement
k8s: 1.16+  
certmanager: 1.0.4+  
kustomize: v3.8.6  


## Deploy

```shell
make install 
make deploy IMG=lut777/pingcap-webhook-test
```


