package main

import (
        "flag"
        "fmt"
)

var (
        cachedEntries []int
        useCache = flag.Bool("cache", false, "Use cache?")
)

func cachedFib(x int) int {
        if x < len(cachedEntries) {
                return cachedEntries[x]
        }
        f := fib(x)
        cachedEntries = append(cachedEntries, f)
        return f
}

func fib(x int) int {
        if x == 0 || x == 1 {
                return 1
        }
        if *useCache {
                return cachedFib(x-1) + cachedFib(x-2)
        }
        return fib(x-1) + fib(x-2)
}

func main() {
        flag.Parse()
        if *useCache {
                fmt.Println("Cache has been enabled.")
        }
        for x := 0; x <= 40; x++ {
                fmt.Printf("%d ", fib(x))
        }
        fmt.Println()
}
