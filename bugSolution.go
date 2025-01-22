```go
package main

import "fmt"

func main() {
    var m = make(map[int]int)
    m[1] = 10

    value, ok := m[2]
    if ok {
        fmt.Println("Value for key 2:", value)
    } else {
        fmt.Println("Key 2 not found")
    }

    value, ok = m[100]
    if ok {
        fmt.Println("Value for key 100:", value)
    } else {
        fmt.Println("Key 100 not found")
    }

    delete(m, 1)

    value, ok = m[1]
    if ok {
        fmt.Println("Value for key 1:", value)
    } else {
        fmt.Println("Key 1 not found")
    }
}
```