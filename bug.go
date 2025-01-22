```go
func main() {
    var m = make(map[int]int)
    m[1] = 10
    fmt.Println(m[2]) // This will print 0, not crash
    fmt.Println(m[100]) // This will also print 0, not crash
    delete(m, 1)
    fmt.Println(m[1]) // This will print 0, not crash
}
```