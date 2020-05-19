package main

import (
"fmt"
"net"
"bytes"
"log"
"os/exec"
"strings"
)

func ececute(var command string){
	cmd := exec.Command(strigs.Split(command," "))
	return cmd.CombinedOutput()
}

func main(){


}
