package main

import (
	"context"
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/filters"
	"flag"
	"fmt"
	"os"
	"os/user"
	"os/exec"
	"path/filepath"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [mountpoint]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func doall(mountpath string){
	_, containername := filepath.Split(mountpath)


	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if containername[0:1] != "/"  {
		containername = "/" + containername;
	}

	f := filters.NewArgs()
	f.Add("name", containername)

	ctx := context.Background();
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All:true, Filters:f})

	//for _, container := range containers {
	//	fmt.Printf("wtf %s\n", container.ID)//.State.Pid);
	//}

	if err != nil {panic(err);}

	if (len(containers) == 0) {return;}

	container := containers[0];

	containerJson, err := cli.ContainerInspect(ctx, container.ID)
	if err != nil {panic(err);}

	user, err := user.Current()
	if err != nil {panic(err)}

	root_dir := fmt.Sprintf("/proc/%d/root", containerJson.State.Pid)

	scmd := fmt.Sprintf("sudo bindfs -o force-user=%[1]s,perms=0777 --map=root/%[1]s %[2]s %[3]s ",
									user.Username, root_dir, mountpath)

	exec.Command("sh","-c", scmd).Output()

	//fmt.Print(out);
}


func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage();
	}
	fmt.Printf("mounting %s\n", args[0]);
    doall(args[0])
	////Tree command
	//mapped := groupContainerToNetworks()
	////showTree(mapped)
	//
}