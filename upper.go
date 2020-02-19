package main

import (
 "bufio"
 "os"
 "strings"
 "fmt"
)

func main() {
 scanner := bufio.NewScanner(os.Stdin)
 for scanner.Scan() {
  words := strings.Split(scanner.Text(), " ")
  m := make(map[string]int) 
     for _, word := range words { 
        m[word] += 1 
    } 
    fmt.Println(m)
 }
}