
package main 

import (
    "bufio"
    "os"
    "log"
    "fmt"

    g2 "github.com/AutoLogicTechnology/Gate2/gate"
)

// Simple CLI tool for generating a new QR code
func main () {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter UserId: ")
    userid, _ := reader.ReadString('\n')
    userid = userid[:len(userid)-1]

    g := g2.NewGateAndQRCode(userid)
    g.WritePngToFile("my_qr_code.png")

    log.Printf("User Secret: %s", g.UserSecret)

    for {
        fmt.Print("Enter Current Code: ")
        currentcode, _ := reader.ReadString('\n')
        currentcode = currentcode[:len(currentcode)-1]

        r, err := g.CheckCode(currentcode)

        if err != nil {
            fmt.Printf("Error checking code: %s\n", err)
        }

        if r {
            fmt.Printf("Correct: %s\n", currentcode)
        } else {
            fmt.Printf("Wrong: %s\n", currentcode)
        }
    }
}