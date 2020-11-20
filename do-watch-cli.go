// do-watch-cli
//
// Author: Ondrej Sika <ondrej@ondrejsika.com>
// Source: https://github.com/ondrejsika/do-watch-cli
//

package main

import (
	"fmt"

	"github.com/go-cmd/cmd"
)

func run(name string, args ...string) (stdout, stderr []string, err error) {
	c := cmd.NewCmd(name, args...)
	s := <-c.Start()
	stdout = s.Stdout
	stderr = s.Stderr
	return
}

func runStdOut(name string, args ...string) {
	stdout, _, _ := run(name, args...)
	for _, out := range stdout {
		fmt.Println(out)
	}
}

func main() {
	fmt.Println("=== do-watch-cli ===")
	fmt.Println("")
	fmt.Println("List all resources from Digital Ocean")
	fmt.Println("")
	fmt.Println("--- Droplets ---")
	fmt.Println("")
	runStdOut("doctl", "compute", "droplet", "ls", "--format", "Name,VCPUs,Memory,Disk,ID")
	fmt.Println("")
	fmt.Println("--- Volumes ---")
	runStdOut("doctl", "compute", "volume", "list", "--format", "Name,Size,DropletIDs,ID")
	fmt.Println("")
	fmt.Println("--- Load Balancers ---")
	runStdOut("doctl", "compute", "load-balancer", "list", "--format", "Name,ID")
	fmt.Println("")
	fmt.Println("--- Kubernetes ---")
	fmt.Println("")
	runStdOut("doctl", "kubernetes", "cluster", "list", "--format", "Name,ID")
	fmt.Println("")
}
