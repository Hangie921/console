package main

import (
    "log"
    "serial"
    "fmt"
    "time"
    "os"
    "strings"
    "bufio"
)


func main() {
    c := &serial.Config{
        Name: "/dev/cu.usbserial", 
        Baud: 115200 ,
    }
    s, _ := serial.OpenPort(c)
    
    defer s.Close()

    read(s)
}

func write(s *serial.Port) {
    time.Sleep(time.Second * 15)
    fmt.Println("start to write")

    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Shell test")
    fmt.Println("---------------------")

    for {
        fmt.Print("-> ")
        text, _ := reader.ReadString('\n')
        // convert CRLF to LF
        text = strings.Replace(text, "\n", "", -1)

        if strings.Compare("hi", text) == 0 {
            fmt.Println("Hello, moto")
        }

    }

}
func read(s *serial.Port) {
    fmt.Println("start to read")
    tempString := ""
    reader := bufio.NewReader(os.Stdin)
    for {
        buf := make([]byte, 128)
        n, err := s.Read(buf)
        if err != nil {
            log.Fatal(err)
        }

        // parse the data
        switch tempString {
        case  "Username:", "Password:":
            log.Println(tempString)
            text, _ := reader.ReadString('\n')
            text = strings.Replace(text, "\n", "", -1)
            s.Write([]byte(text))
            // enter = {0xD}
            s.Write([]byte{0xD})
            tempString = ""

        default:
            if string(buf[:n]) != "\n" {
                tempString += string(buf[:n])
            } else {
                log.Println(tempString)
                tempString = ""
            }
        }
        
        
    }
}


