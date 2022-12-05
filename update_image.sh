#!/bin/bash
read -p '请选择命名空间：' kubenum
mkdir -p /opt/set_imageg
> /opt/set_imageg/kube_name.txt
> /opt/set_imageg/kube_images.txt
> /opt/set_imageg/kube.sh
mkdir -p /opt/set_imageg
kubectl get deployments.apps -n ${kubenum}|awk 'NR>1{print $1}' >>  /opt/set_imageg/kube_name.txt
kubectl -n ${kubenum} describe deploy|egrep "Image:"|awk '{print $2}' >> /opt/set_imageg/kube_images.txt
echo 'echo kubectl -n spaas set image deploy $1 *=$2 >> /opt/set_imageg/kube.sh' > /opt/set_imageg/yang.sh
paste -d " " /opt/set_imageg/kube_name.txt /opt/set_imageg/kube_images.txt > /opt/set_imageg/kube_paste.txt
while read -r line;
do
  bash /opt/set_imageg/yang.sh $line
done < /opt/set_imageg/kube_paste.txt

cat /opt/set_imageg/kube.sh