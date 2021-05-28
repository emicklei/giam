package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var verbose = flag.Bool("v", false, "verbose logging")

func main() {
	flag.Parse()
	snapshot := loadMemberWithRoleInProjects()

	if "roles" == os.Args[1] {
		//temp
		member := os.Args[2]
		for _, each := range snapshot {
			if strings.HasSuffix(each.Member, member) {
				fmt.Println(each.Member, each.ProjectName, each.Role)
			}
		}
	}

	if "owners" == os.Args[1] {
		for _, each := range snapshot {
			if each.Role == "roles/owner" {
				fmt.Println(each.Member, each.ProjectName, each.Role)
			}
		}
	}
}
