package main

import "fmt"

func split(sum int) (x,y int) {
    x = sum * 4 /9
    y = sum - x
    return
}

func split2(sum int) (x,y int) {
    x = sum * 4 /9
    y = sum -x
    return x,y
}

func main() {
    fmt.Println(split(17))
    fmt.Println(split2(17))
}
