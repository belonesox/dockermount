# DockerMount

> DockerMount gives you ability to mount internal content of any docker image 
to given directory. 

Just provide any empty directory, where last part of path equals to container name.
and run  ''
```bash
dockermount /mnt/dockerimages/gifted_fermi
```
and all guts of container with name «gifted_fermi» will be mounted to /mnt/dockerimages/

## Prerequisites
- bindfs > 1.13.10

## Build binaries and install
## Prerequisites
- golang/dep - dependency manager 

### Build
1. Resolve dependencies 
```bash
dep ensure
```
2. Build binary
```bash
./compile
```

### Install
1. Move binary to /usr/bin and change 
```bash
sudo mv dockermount /usr/bin/
```

## License
Apache Public License