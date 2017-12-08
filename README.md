# aced-dns

* kubectl exec -it alpine-85cc59f6b6-8lmvf -- /bin/sh
* kubectl exec -it aced-dns-6687589f4d-ss8zz -c aced-dns-container -- named-checkconf -p
* kubectl --kubeconfig=$(pwd)/t17b.t17b-zone.us-west-2.nonprod.cnqr.delivery -n containerhosting exec -it aced-dns-108325911-n9d6m -c aced-dns-container -- named-checkconf -p

