### GO installation steps for linux ###

1.  visit the official go downloads page and choose the latest stable go version
```
https://go.dev/dl/
```
2.   download the go tar
```
curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
```
3.   extract the tar in /usr/local
```
sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
```
4.  add the go path to **PATH** environment variable  in **/etc/environment** file
```
PATH=<existing path>:/usr/local/go/bin
```
5.   source the environment file 
```
source /etc/environment
```
5.  test go installation
```
go version
```