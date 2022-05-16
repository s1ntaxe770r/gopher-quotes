## bootstrap a civo cluster with nginx pre installed 
cluster:
  civo kubernetes create consul-cluster --size "g3.k3s.small" --nodes 3 --wait --save --merge --region NYC1 -a=Nginx --create-firewall="80;443" 
  kubectl config use-context consul-cluster

# install consul using the consul-k8s CLI 
consul:
    consul-k8s install -set connectInject.enabled=true -set controller.enabled=true -set gloabl.datacenter=gophercenter1

