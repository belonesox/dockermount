# DockerMount

> DockerMount gives you ability to mount internal content of any docker image 
to given directory. 

Just provide any empty directory, where last part of path equals to container name.
and run  ''
```bash
dockermount /mnt/dockerimages/gifted_fermi
```
and all guts of container with name «gifted_fermi» will be mounted to /mnt/dockerimages/

And you can edit all files using IDE/Text editors with GUI, run git and other deploy operations 
from shell of host box.

## Prerequisites
- bindfs > 1.13.10

## Build binaries and install
### Prerequisites
- golang/dep - dependency manager 

### Build
```bash
./compile
```

### Install
```bash
sudo mv dockermount /usr/bin/
```

## License
Apache Public License