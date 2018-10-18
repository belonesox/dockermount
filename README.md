# DockerMount

> DockerMount gives you ability to mount internal content of any docker image 
to given directory. 

Just provide any empty directory, where last part of path equals to container name.
and run  ''
```bash
dockermount /mnt/dockerimages/gifted_fermi
```
and all guts of container with name ``gifted_fermi`` will be mounted to ``/mnt/dockerimages/gifted_fermi``

And you can edit all files using IDE/Text editors with GUI, run git and other deploy operations 
from shell of host box.

## Prerequisites
- bindfs > 1.13.10

## Install
```bash
sudo mv dockermount /usr/bin/dockermount
sudo chown root /usr/bin/dockermount
sudo chmod a+s /usr/bin/dockermount
```

## Build binary from source.
Not needed if you download binary.

### Prerequisites
- golang/dep - dependency manager 

### Build
```bash
./compile
```

## License
Apache Public License