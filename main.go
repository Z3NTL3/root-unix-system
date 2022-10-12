package main

import (
	"fmt"
	sys "golang.org/x/sys/unix"
)

type rootMe interface {
	setRoot() (int,error)
}

type uid struct {
	val int
}

func(i uid) setRoot() (int,error){
	err := sys.Setuid(i.val)
	if(err != nil){
		return -1, err
	}
	return 1, nil
}

func main() {
	var root int = 0; // Root is UID 0
	systemRooter := uid{val:root}
	
	_, err := systemRooter.setRoot()
	if(err != nil){
		fmt.Println(err)
		return
	}
	fmt.Println("You are now root on the system!")
}