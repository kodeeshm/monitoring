package main

import (
    "fmt"
    "os/exec"
)

func main() {
    //Hostname Information
	out, err := exec.Command("hostnamectl").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Host information collected...")
    output := string(out[:])
    fmt.Println(output)
	
	//Dropped packet details
	out, err = exec.Command("netstat", "-s", "|grep", "-i", "dropped").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Dropped packets details...")
    output = string(out[:])
    fmt.Println(output)
	
	//Tx and Rx packets details
	out, err = exec.Command("ip", "-s", "link").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Tx and Rx packets details...")
    output = string(out[:])
    fmt.Println(output)
	
	//System Uptime information
	out, err = exec.Command("uptime").Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println("Uptime Information...")
    output = string(out[:])
    fmt.Println(output)
}