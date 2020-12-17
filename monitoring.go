package main

import (
    "fmt"
    "os/exec"
    "bufio"
    "os"
    "time"
)

func check_err(err error) {
    if err != nil {
        fmt.Printf("%s", err)
        // panic(err)
    }
}

func main() {

    // frame the output filename
    out, err := exec.Command("hostname").Output()
    check_err(err)

    currentTime := time.Now()
    current_date := currentTime.Format("01_02_2006")
    epoch_secs := currentTime.Unix()

    output_filename := out + "_" + current_date + "_" + epoach_secs + ".log"

    // create output log file if it's not exist
    file, err := os.OpenFile(output_filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("File does not exists or cannot be created")
        os.Exit(1)
    }
    defer file.Close()

    w := bufio.NewWriter(file)

    //Hostname Information
    out, err = exec.Command("hostnamectl").Output()
    check_err(err)

    fmt.Fprintf(w, "%s\n", "Host information collected...")
    w.Flush()
    output := string(out[:])
    fmt.Println(output)
    fmt.Fprintf(w, "%s\n", output)
    w.Flush()

    //Dropped packet details
    cmd := "netstat -s | grep -i dropped"
    out, err = exec.Command("bash","-c",cmd).Output()
    check_err(err)

    fmt.Fprintf(w, "%s\n", "Dropped packets details...")
    output = string(out[:])
    fmt.Println(output)
    fmt.Fprintf(w, "%s\n", output)
    w.Flush()

    //Tx and Rx packets details
    out, err = exec.Command("ip", "-s", "link").Output()
    check_err(err)

    fmt.Fprintf(w, "%s\n", "Tx and Rx packets details...")
    output = string(out[:])
    fmt.Println(output)
    fmt.Fprintf(w, "%s\n", output)
    w.Flush()

    //System Uptime information
    out, err = exec.Command("uptime").Output()
    check_err(err)

    fmt.Fprintf(w, "%s\n", "Uptime Information...")
    output = string(out[:])
    fmt.Println(output)
    fmt.Fprintf(w, "%s\n", output)
    w.Flush()
}

