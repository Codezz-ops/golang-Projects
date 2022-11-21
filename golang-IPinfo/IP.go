package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "strings"
)

func convertToBinary(s string) string {
    res := ""
    for _, c := range s {
        res = fmt.Sprintf("%s%.8b", res, c)
    } 
    return res
}

func main() {
    ipv4Lo := net.ParseIP(strings.Join(os.Args[1:], "."))
    if len(os.Args) != 5 {
        fmt.Println("You need to pass 4 numbers")
        os.Exit(1)
    }

    for _, i := range os.Args[1:] {
        number, err := strconv.Atoi(i)
        if err != nil {
            fmt.Println("Failed to parse number:", number)
            os.Exit(1)
        }

        if number > 255 || number < 0 {
            fmt.Println("Not a valid input:", number)
            os.Exit(1)
        }
    }
    fmt.Println("IP: ", ipv4Lo)
    fmt.Println("Binary: ", convertToBinary(ipv4Lo.String()))
    fmt.Println("is loopback: ", ipv4Lo.IsLoopback())
    fmt.Println("Default mask: ", ipv4Lo.DefaultMask())
}