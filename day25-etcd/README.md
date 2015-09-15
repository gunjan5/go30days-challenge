curl -L  https://github.com/coreos/etcd/releases/download/v2.2.0/etcd-v2.2.0-darwin-amd64.zip -o etcd-v2.2.0-darwin-amd64.zip
unzip etcd-v2.2.0-darwin-amd64.zip
cd etcd-v2.2.0-darwin-amd64
./etcd

./etcdctl set mykey "this is awesome"
./etcdctl get mykey
