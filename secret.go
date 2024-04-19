package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Prompt the user to enter their text
    fmt.Print("Enter your text: ")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)

    // Save the text to a file
    err := os.WriteFile("secret.txt", []byte(text), 0600)
    if err != nil {
        fmt.Println("Error writing file:", err)
        os.Exit(1)
    }

    // Exit the process
    os.Exit(0)
}
