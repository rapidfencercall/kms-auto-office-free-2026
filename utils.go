package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.0" )

func oKfXjGF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qf5tYvJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PWXLjgflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1JlVhaOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BLbmqrsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H8Ufp3niWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gSyUG4V4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func neOeOWy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cCBlxmyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GSP6noyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sAfE2PxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TfqYhPuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bVwfwOgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eHWJKcUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNR9EvrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sXrKj7zBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ooU00vqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pF7A1WRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Q8SzEAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TMFVuXaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SEG33hFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AVKaDp41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GjwtSDHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W2bdRlASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XqhGQXNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cVRZ2rm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEKC5gv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DdFJDfOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fKZuPqrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FCZUJYY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SKiIwEt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a3fwwj2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xVlpZU5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gKgldAOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2HhfwQJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func atxXBl0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R0PVPNNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NjSwlJYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4XA2e3UyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HMy8l33UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0XoSK2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V13BXva5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hww3KkuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M55Js0pgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QIA81qLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func up82pS6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2Cv0oBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LeDQ3yvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2pNEFLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ESKJaQzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ps2L9VaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Racuvg2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SafH0nj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5iaKvV4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vrPUaocYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PpDNRFUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DK8NJckpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AKcFVd4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 53TexVtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3srK1LTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FbZqBO3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LAu9G7jEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jhqSevakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nUv6DEkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ea5aT2V8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YaJmJl6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X1UCCnm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYisAqQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1EofhtmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a7OPe44iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J0XEYCRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2cT23K8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YD2pXCKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1gCF91PvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQReeLpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CiLZn4s8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xr8V6X4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJ3IZAb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GaSym1HIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W0wgaIIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VeuotwvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wXwXONJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nKHKrZXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5e6u2d9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gazuo8YmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VwyXnEKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sd7Ve6H4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EdRW0J23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FpwxMv7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YKCIt1UoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cn2Ha7lYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func beDcE8GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RVm2i3clWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wr3qjuPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func huQSKB8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Itb4KuDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l03tGSgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wtseIgYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CoZ7WP8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdRavsmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OoaIq4RaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fBIGNELkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1rco5hCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1IYTycWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gEj7SSmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MviYTdWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IcNYV4KJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B5FjhEX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7dDuwX6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yQtjDqNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kwyvW8rUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BfqSsdvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BnX2ACQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WtHjUiIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GyuIaIqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KgNF77sfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 26l4qLTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HgcIVvLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ylXOABMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mopDJOojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sAHnmpiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pga3Pxw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tG2ulWl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PPGjMZRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0rOvF7vSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VHeuLErUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4PTWMh4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eD36OLgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOk0CRzaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rib1WW4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSf1mWFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DRkIQX3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VmsPouAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bLbrNjk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JtduGWsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iYLbihlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gd2Snz9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9LgImNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CFGzqwpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eQ2mpZIxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f13Uh5W2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y6AzdnqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eP0tG4qXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XqcoByktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LXDIjd6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L0uwBjybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ZA2yAknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a7EfV6qAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NwGQJrA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eWsAQC6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c3xYJZQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AtHr1nl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QcOKgsKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6wCGn1QJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ra4VVBxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UwMa4bcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cmjOY9u7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bmBQcdGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aUrLJvavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func klTL7qzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PXO0eT3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EplGoMq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8LAnMqR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oqO97rJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PEA5mqsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h9XD48sbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ycaZ7oTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChtbImsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCrUT8Z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DNszwPlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CQkkKDcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxjXNBOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func byPhOnVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C3OJZMlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B0hieNg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucGFyqNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zJKPnZUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3QYsOqBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hgXYMgOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AcgQ1E7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 468ACqpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9cwiWpbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VP4v3Q2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MwYclbtzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9UMZw4irWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H1Kie1AaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z0UJc2JVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NUCfEyIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7uqhZj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4kzUJKdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LL9sm08lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jb6mwhgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jOjuJ3GVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2h5rdW2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SOkDwZqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PpMxiwn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qtyIiDkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b1sy22oZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EGOLy1VdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nn1d0uFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2ggLVqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 08RjArrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FEMOJSF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ifEMvPLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8wrsLhU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nDRDm0dhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hJ2b1L3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T9tsuWxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VyyAbHeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gyfoi7QAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nGvvZQBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cai0ERmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O1HLdp13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qTBTkUeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1gNl9NaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LHFhLwCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w1yTOrJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLMwCiuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ta8JwUD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UvJer5f6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YYSq4DqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OscoHZQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b3HXxLndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NDmpDNwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ApTKh3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gn8UbsQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2qMgspOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VY8Qhl7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LtiddRryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yotlAdBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OuopfmisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBzNiWCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ODoOs28uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HTaGDWZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KZoEAfSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3UMMhHF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yGYIdQehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E1ePoehsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L8CeyrJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tUQV6JDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aq6gPQboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rwxxhz5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8XurUpgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func il3AZTyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wkmhAfDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func He542hYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7VYCQ7V7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WifecECeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j5WU6HMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7KR7bnLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func epbHK91vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JKgU1CcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lphgHuo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lx7dhfY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LM33oQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s4Y6mvJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e2ymxgmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TgzqXU6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rnh2EwjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JVcEwHrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0pCxZ3REWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hJJoiQNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func arodUABEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c30ArA0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S7sDpn02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dEEGVLnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nw38EIZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w2CU3jrgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opqVJI9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ef84LSHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9jhlNGYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WNAljT2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qphQkLAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MCsgcdOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S5A8mQIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dR8Jwva3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 042DUAT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HicfEduXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0c5xCDEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cg8k0oXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cEnfUiZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbKGQiOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func naRnK3ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1OQqruWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8q7KY9ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bOWNmu50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xYKAnn6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v5CW1hvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3kqBrFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0p7pNW3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0uMBKMfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hU1mIK2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GHDRgdNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8hs2jb3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XinugXt5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H5HpOaxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1sqUDNpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qoZgjXadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TizPdJQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6yrQWN8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HnbTvOt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zLGBavaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i12assZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NgJSiKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3RHlrwefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TgPFIBmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cg86pVobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0OijTzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9dk7vJJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p2AgxiK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4xdKT9gRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PivzRy1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQyzJmTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CkVwxKwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C5UbZ2LuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8i1OQ9MlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g4F09NmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MKuz1IGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H5FE9pHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w832B5oFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dww31aTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uEOYOdbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xL3OGUz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mAUXNBnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lDWTu6tNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PPmNtWCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G3y6oCMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QS977IUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YvoDx0NgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O6RBPZd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44sbDZS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4hM3OUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gqYvMcAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9rdZuHyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VCeORK9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NikTSG5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZEJhmNMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sD1yHYh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tJeVLADmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PtpmszpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFcITxlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ND3zREFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func urIkwLC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HBEwULl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJ6zCfSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FqiDipXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEYx4bxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qOReql0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tV916huWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oGeJeKoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6qpKlTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NX9Ci4vXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2MQdPu6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0YyLagijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AKqRSIaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RLnKyQNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e2qGHiSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K3kjTw8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DcDvvfG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3in9QNp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iQSNqrQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RDuz9b3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uk5wE9TaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k87TGu2lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VpHOviQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K8pnGkQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0wQt2O6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q9LeTuDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VekZUzTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tXytW7i6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func igxKulCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wsp5Z96WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CTzb5x8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5U4iUOvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yBeK1pRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nSy7JZtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AW73iRB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M0tj3jxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func moRjJyezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bzmwbhCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZSD72954Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Moixk30MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B3impMFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gRT7GoJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AfTmgQevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tv5ikyKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5gyRV7G8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SlBoBjjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THA89aMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oBissw22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0UWyYG2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wUJ36pFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0uZuLDoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1p7RvlDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VgnNtmiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vdJy5nXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h8srmzMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEtiLLBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kogJ31IOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nMAIUHhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKhbHe5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tDDdIjFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9PhSnenTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R9b0A8V4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4b33YIjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1RHYAs0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cqsKzrWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ha0m0qXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 48Eb2V7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8iFpOaNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SiRkLUaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S85uBifbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4lzeil4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2W4sJoeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func El6zvgtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85NBlai8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eOawOP62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sZlI2ycgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kBqQVOMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bYTMaKeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SxUeavNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oh2BQfzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mZq8o5byWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z4PF14FWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1LHqhKfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aatA8bCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QjiJDMCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8lmplwgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J3p9xgAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KMnVdawqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aJqrckAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SZfb602oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yrkCDTW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nBioOOJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NE7eGkQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p7Q4BR8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MW8BAKnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HERYPYVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sCCDTjylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lPlvCaoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eGEDeElIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UBGzM0CWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 40U7q4aBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PhTmD1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dSil9XnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mFugOrBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vSl3MgqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Wc24wqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F14eRovcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ApDYUxOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqtKKfKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLrBWV4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ODk4Cn57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3e4rJY8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0OhbI5v0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NH39NgsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q0yLpwxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IVnPMcnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bSEovtvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oHulpKw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jpUlVq0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y3E3UWxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func svTW9g0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iaBtA9kUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wmkJwkxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BamEyrABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sh2KXKmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vJsgXqcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jE6PlLUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zeH4X5uWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V2gW0ABaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iNgLL2mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qIhyJTMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 60lbem0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SZONAJ96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BaxYtdgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aIz1CI2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bk7CHTc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wlRvQKukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tyo19H51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvItO14PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func er0Ct1bpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ExRSFIOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i9P99s7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rDhT0qENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ditkfJSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U2T7qMxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9CgAirSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mzwXurD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TRgEo58RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qT9cZv0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GckU8mggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HoFSc6OuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfeGyTx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iRZ64r5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func czCq6PDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yXw0qZRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lq6OExH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eoz1W0ZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pDZK4OiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YvalQkNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h2NRb0WOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJGzEZbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LPXvHtqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h1OlpCEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9bfxQpAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ckxkq4guWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L20lpaiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Ps4XGwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YaminnTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aZ9KjPUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ouJ7ItWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H7SqBuXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xB46qkCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N3avXl2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tn56Efl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1mjshcSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 98igK2jcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lgNFlZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8TqYPrX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6QHck8IpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEX9dSx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UUCFx8tcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UIVGsPisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BpnQKzaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oTEX1gT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LLZWXYjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qnEsOYxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmWCmcpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1l6jxyFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q4rZyZeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qIW3zOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VSiKzCjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fFfoyES2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hwIdknlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fHpTDQ7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GPCDLycfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rs7pa8dWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func onpv3se0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpXHmbFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4SrbbRZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9KYiMNkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T8A2MmaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvERmvsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HFuXMwlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtgk6yWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func clyZ1AZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8YKbzHJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VU2ZVkv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uYqqwuTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVSpHKArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YfCquXlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lCGhu7syWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rdC7awi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Ben8UkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uNQqeyTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LtBrWBlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ZUd7VqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r1ChbzOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N2yNjsMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QyCJBWooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uqWelj7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u30dh6CTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b97efarDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBdRUlFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8t8i0BvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iCTugeihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aUkJeLUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvmNLY0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func giy1S4lkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ybmn0CLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kQi6zohSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RmsqxcftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8bbv28gkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O7DLlvJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func epsqkD2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dhl3by6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iI39tdHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A9bll1OuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jVUNm5TLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0M0aoyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tw2Z2v2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GmymxYahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zjbER3l9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nMteqEL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2AIiOhiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2CVaDgdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZrijkzmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8SZc8cbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JnFjhhL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qopIfyTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rLmAZUWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JeEs68WPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HY4n1PT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zLigWiz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v9WtwtyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XJY9Ka77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BzdvTU9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztoGNYhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e4nigO18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WheEAR3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qjZ427HpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsQNq57bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y5iB7pbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SKXeITqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nhiakTOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0J8lRlcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ze0O5OKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKciuRbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cVstuLc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6lZ1OfpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YUNpDGhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cmXVi5tJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OY0mBQvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t99DeaSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cyK390lHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WNyEE119Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KP4R5LzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1dNVOV2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mcz5HJBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uKKBl7qpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oTYTrQTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ft92EsQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yp97Jh9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PL0FikrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 44f2xj8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xBZs1uw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h098w1CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L9DowmXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OVCZ9xY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C1av0SaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7tePoPXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rj3hqbf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNANXL3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b41VWzSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AztAvTkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UM0PJDc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SNUDH3YcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aqLFbHRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yGWvRf8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LduK7ztFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tf99SlioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 99XpdPqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ODTKdJqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func volwEMavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ha2GBGPxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8gkOTTaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zcr58bYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gwwi4tUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tpy6kLvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YrsxFEwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1AMXLfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqsnAqasWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 43AKQj6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cIE2jFklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RHSpQ5lrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YD4x0loHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jz3Ckeo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t4wYMRyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x04cN4oPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TXjK9HSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i3MiNcBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R30C6oBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func twJsdnxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nD1KqaV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VV5LP8InWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M4n58zIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oWKHgMJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLRiQ07HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GYYZemoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OcN7SEaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z64qcsBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lxFI5C38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OFID6GkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lGHFGpA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IyxriHWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CyYnemMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBRil3nSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func poioLRW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hDhRBdU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QM6y87o7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kcQoFLfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bBV3an0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EaggFE9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PjWNIEScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uIIx6jxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3HYZ5TMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzgpnRN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WtCO0nqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RAFUbl6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func auioDV7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qUTCemZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EqFVzYtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gPqBxDbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C20oT2THWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2GzquIn1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func db0qrBa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ocTQxxxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j4iazw5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cHEGVTt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CAoaPoChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qkV4uzFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W03nEoZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0I58E5NmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iVRjYjirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 35WlNHx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8pCMTCnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5k2EHKRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tec2pNKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DWgFlVuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PtBzrjLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TTJrbvitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zryMG2LCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JxNsAyxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d3yAZhl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lcwlBuxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8dDO211wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8E5Iu8KuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J7p9U3SFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z2HkgQZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JbO63soDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lWnekhwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zE9c0M5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func taXDD8mGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xwgoz0eWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bFb9osR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func neLmF0TFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pWZqfh0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XjB1uqpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dh86QJoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P2zi2xtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VvIfb8JuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E1se2hHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func goVNZrGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bSett6MfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hD6p306iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9V7kwrcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qLHmeV23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qj1el36XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V0qpKGxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JlIyPrLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mit5OQTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WjeOTgIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChcKyn7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QPXOkLrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4DVQFlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IAAKKNzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGzYR7ZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lXxwTbTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mko8qcaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OVwp2dfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HlgTaGaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 68U4A74GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygH6A0wKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XYeJwIC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7hTKldsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8xlcALpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jwcKxflMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cTaUuKHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2wcBdmaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yUdZgHfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 42rIIAnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZzNU1TwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lggSeZlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xeU2x2oxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2pDOkU54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mDaNEMKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PsOVNUeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fRFPpMGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7UmBKdObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K2XrDZ6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OyHqwXznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LOm5yDX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQ4Bxc9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YpQDA89wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1jy1cwm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FeuHfpiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qCfCvz1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yZ0HJTztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aGlEsDtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TJp939EFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z14wMMMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k8AgEBzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9PTKp0D6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mzUPwjIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HaU4sZhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YgImf7nMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ECsCXJGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EushUvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LVshPTt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func te6fzFwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8cAZQxfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GthY24pbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ukS3McTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xrpKUFmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W023kMFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pDvRN85nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d8nSa0WcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T8u4GGlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0g18IwoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GQM3v6FnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PMidoUQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6SPneb8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SvZ0O2ihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aHHnioGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23QQiXAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mvozfuUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 96wbdgA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FVu8QBQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKLabtiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yT42dTiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLKoyAZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 24WwSwHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jntv4QMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ty4hz5npWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJeFoBd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmZYp6XcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zS4xtqBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kO0w7XDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b34hOiyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WPIJg5dgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ArP3P4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RK8lRPfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AQVO2qlLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zAkKw5U5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4folVrf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z4Gqs6VAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 351RAUwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func efcl7SsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tfWPLnmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IWvYKcuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VuJ8EmDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMXOoIEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5V0q0ksLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NK6QvH0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rtoaz8H2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xzOVUk2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QfF68x6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sAJWXtSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PcueaUZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SWz4zqgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tsll6DyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UBnN7WdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7rh0sK68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kqV9mT7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uSGzaMbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3FxjGrQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kiYzoCGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wdX6S1zDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IBinOIp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GVywV5tFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ey7bJhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3h9pl4gjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZIvJwcUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PAlGa01HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xhSL8bVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6shFnBMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AOQruhCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func atmJj77kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CPcX7uH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KtNiqK8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kIgHsQ0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WlCM8rKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bcQpzhjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WpGXsQW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VOb7k8BeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0jHwSKTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GKhGR7JeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1w1EvbB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nW9A140sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func haYcDISMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H66OjOsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XWvheiTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q3onAc0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qk0CkeInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 82DBMwAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aTdke7FPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kh1bVnGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THx78gFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjJdsZHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NB9f9SQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LriKw4xRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FBGmVTejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6s7rxsS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2COdgcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Rwy3jY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cc7Dgi7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OsWlemzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sV2Y5XnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z5FdlSA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eVYaFXkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NUslmOwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nc9D9VoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w8y7c9lBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FfJRGdqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FhtSr0EAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Th9g7pOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gXoFF6lIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FNWBP21QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func umLu3Ac7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUWGsjHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func THZ5ns1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P5tinITEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S6YoDuOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LJxEx3Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjKdl0NsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MFBpsoj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBPlCd7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bR9YvmcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzzGHtFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QJ2RvQXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5LzBQeOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XilgoN50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfUuwvADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A0UfwifAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dWICtn1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func te9Uio6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A6Ku5ZOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OoKysapmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NFRm0N4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NBImdpRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OJ45Tr8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IoqyqFTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AgwfubGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KPKkIavuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvPzWZM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7DzFMzXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u4XA2HrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GQ6Z0EKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QcjAa074Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UWPm16YAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D0r8mxWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e1KmzeFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8kLPk7HMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4PnbTK9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gXQO21swWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6NGnx8hAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yEzLEsVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XwvF5PDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5gBd114Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XilspY7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDZX41KaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G7jU25nxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QYwBlN0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func stEO3x1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D3dI5rwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TvCgK5d6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y3uTprktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KfWDttZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2dA7SvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e2SMT6xwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IsUGWgIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CHoS3IbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kc65K3cJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8gaod1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 31w111MHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gOMywKuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p02KNaXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Up3eLjrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qwu52OtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4MOBKONiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gx8M3YR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pzCkkkyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5iAeN86yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MEKGxb8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ONYNswzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lLo62cnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D18XtV0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 41GK3SEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S8CFW0inWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t0xYLsPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MNQVMNZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wrJywsULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func biUPko9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHnb5qnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VmSCSFeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zwFRwHi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AxttFsw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dCXz4UNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fXklMWPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FqBT6TWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0koeLe9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jhrhDfmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bpOU9LFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXyhXPaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rObJDQP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B8fWF29LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wwXSTQ3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VAgTLHb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DtQioBLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gPReVc6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4U8yxMjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n9ETsd6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RXq3wnHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w8WWpX84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ltOObqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oVd1G0zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HBp1M1BMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 25mAjMnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OcaDbfigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MM2fQBbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vEencyZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7BVHZk3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Hj2ZsmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zh4NREMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M6gHe9S8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ips31vMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Id2HcZLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hj2SsY5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kxau7cnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CjFwcGjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k8B4BNUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ta9caP0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qBkzzfpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6MSVMD19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Ra0Gi0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gs6qzRZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yImCvBUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JF8ffHY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WsC8nlIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3EhAr2cdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8xV9kEwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U0HMbt8qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BHyf16edWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T8ydoQ5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PCEfeFXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0zBy0WL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3atIBxcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WxlwSQVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LYzMPliGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 46dLs7bLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z7bxNdzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PdZlVy1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eiLURLiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fni80zQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rviIM58RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Um9hvMS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func peHb1mFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5xjcR4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bujJqYCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PDZGOqrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1feqSGDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZxdZTC03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n2nkOWsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func moN5O362Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8o7ifWSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Pi6WFeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OCZu4dNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dPYpohrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CePUTXmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pCXUInJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5rWFovLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MuAvSXKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dzLhPbsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNaVpEHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2zNyZsDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cAmgoIkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o8S5sGs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LWw8aQmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HlSpfOA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 28HR4UdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EnSipQb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RbC0aTijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X9apoSUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VEKNYGjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdKe7TciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9PkN6OPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qenlYwDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LL4QaCDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func umU0DAo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VfcRBtW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8tDvpMvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQGHPfBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CRpd0AErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MtMGbF6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FiV7lyxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cj4dDi2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sOkQixuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EORMJEZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func va6HHF05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZwDqMmVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GgKPNGYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bpMFgYm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TvfLA7t2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Anm4alTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nEHTBiVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 91YP2vTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nqbQOYC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cuSCAVmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzCZmZV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5LPI8nkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9kKe3XTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bqZwrrVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VsxtIufxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b18hOUsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HiVUIZufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JmCZnvQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L6Udo8O9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 94g9LQDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1HQnsvPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func couGrYgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O4DKW6LYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func scLSzFTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bYEf4zEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4UrZXhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YW4pPrvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func foVd9pb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bssxfm8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DO3ZzAUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5bvogZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4WBBF2RdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func upfSuQLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D6YcPm9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k9hpP5p6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZuY3JKovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ah15KRiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sOkSLr03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EXMqM5E3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g5e5L3UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G2o7Sw7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tNhogRJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nyImIyehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TwfQzntpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJgfEOgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wM8UwRVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vkpo9N5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1fZSSURPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func esZNaY4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJhEJRruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zFD2b3cvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3m03qDiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 58SbZEsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tlZF4fM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NudfPxmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RYqQcYKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQu3XDMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RxBr4rcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ySHpluaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CeaXgXPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRf6QpleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K3oUupDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dy93BICAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VADwSPTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eOoCq0tXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yii3bCy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bSsk3BTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vmmS5gYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6avFnrT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVuU9rqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 89s3hW6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CfoYPpMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vl0mQpe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cYo53BMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NujcJqhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w8VwfHCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2jTPnfuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QqPlSpTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0QFcapAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h2Q5cysPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QVRiD3TSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXknnv1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9FChYspCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2UlLWdVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OSsEWinmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NtbeKdOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K6b4C3A2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dryhPzsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fYt5LDQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3EcpAO54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KzcgE04aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S8su5aOXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l9lV0XuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QlHw7zoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ppvISpXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fjb8xQH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gp9kyWlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o3c02r2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGv9bTFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JLdz6422Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SRhfUmDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sukNSCQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MZnpChYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tYU89gTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m3AZaGmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zmFxZaYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kDSMmmUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XENNb7JnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WKMhhMliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bElRmhTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ApGY0q9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OaF1G1zyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KAtG4GosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xRARUrP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDZhxpWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MXe7hMnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDjd9dQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 54VDvjNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func icsnPCdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ELL7LkdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gdXr96ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nxMT2u6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3uJQU70LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WOTEDVecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GSemOiRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T3xltZwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RdEn5EQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZPySruHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aoz7nMv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m1gkECeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nN6mz53kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxkHDJa3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Wtg9J8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lQCVMecbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0y58W0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8KQt3bWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YCjmnD4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lgKkxMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func poBjX4CWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XwkHhJGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o8NEpAXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YEgXf2MqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rkF36A8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8fsopntSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0r347PxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lsjbvjTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHhDnwmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GK5d05p5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uwyoa1fNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5KTOj7y8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7MsQGkvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func trF8pvhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aKGaRHD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zQ2CcUoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FWjBCKntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZgvQxR4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jD6HtqdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zc4EQ4JqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iMnprTdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nfo8lag0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VjMRUZu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wX93PebSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EAVB3VHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func prYrwWUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxg1WMjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I1WACo4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T1etiP09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mrnGbktzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orpIbPPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VPnH9BzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0yDETC9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kcc74eovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e24O6e36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9c4hWYWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hwkxmeloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxBc6Pf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WYTWmceVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2JLHICuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J1OsLVckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sDhXirT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E1JWaRNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9juS3QPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4q6XDVJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mk4jMSCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k8CODRiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6JNNGQcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xwkG7RoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QxiV29joWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uOGaJgy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRVTRwBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gfIoFdCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func olNwWUwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NXyt1bYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ANb9ogOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IaH0O0rgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2RanEMIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pn8BPfkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iH1YIbrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func APu2UcNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tdhizAxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iyLNXf3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IG1lmjkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iay8nNCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1n49qv6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 98xFLpRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nXpAKRLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CbDztH3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7bQBRcavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbOjQGAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WQo9xNrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jHnt7c2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xmqaIcEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ABMllehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h9Hn3MnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xhDNpL5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ufHDxa59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyVCHFLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func npi2W7G4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyfJtorwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rA24xlh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cTVX4FngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x1MUo2LiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2J2YxNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SVtC7UmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jDjI0nm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hlXRDo20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KOpypPlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OWumlzKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u5IPeF0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBqYLPwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bwZ9654QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z5dQLy48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UBbPSPrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hbYOUcfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3a3DzuBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GUU1OKjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lyRnHPtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMri1DfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFc3e0JrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func COjNxMhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PTMWDaCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WdQAs5BKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kHCrrpFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H5FakbqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func umZfGevcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5wLczN72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0D6gMyGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S2yMf43ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sFqITuUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8bXYNmo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNjuphkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0AqT4TI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nz5SfccLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1vqveRPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UVnXICzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3llAoL78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PHcxW9gcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YHY4f7n1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NO3K0SR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4DlTqja4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n35LWidkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7bnO53uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3aPaxHNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wyFXB2KaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oEoeeT3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hLyrvcSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mYarFcHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uNXNAkC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9MNcPOmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fgj2bzs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mRMkQ3ZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oMAHgnhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0dw94xQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Aad8ZAd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JVBFQfZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8b6ERXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ujJvYHD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RBRJntTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5400ojnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hLZbTqjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hUy4IjzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wlzZMoA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8iGHiiadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0sZmOFD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V1Cj57PkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SZS5ixsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ahW4NCmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QvNDe9exWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrpeON7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BBLHGT7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KEGMvzkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mvzKk8Z5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZ1kLTsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yx1WffJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QU4jizM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NicQIzQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vpC9wph8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aYXOiU0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hcbw25dFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LazGaGrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r3KrbbD6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qq5FauSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3oyltQIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mfnB5wLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eGh7x6V2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K0smgWiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SVrVwVoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qAIMyvUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1zD4u3jHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrbu4UQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dWwQW3xOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cQx2uHB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q8PugJsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hkiXr8dRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uFpTKM2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4fEUHKd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mqYgiqnMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JxUQOZ69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XXapRBiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xN3ClL3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K3AjDXWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7CUIRWT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xovwl2OcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvoWIgIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AfrVwy78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PjDLw3epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ql78qgmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tp3KGEIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E8lckRldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lpAOsDsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PU7tqRwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DvtS1XgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZRtiKlfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L6uLJ216Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lXw8j5oaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vt2UfdP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZz9V4SoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Iuh36H0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jtqfIAikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gN5yBLrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NOM5MuhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ykns8uB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UGzOOe5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AeOGBRIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHYDSpdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S7jPbrgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzInSlnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yNzifFqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HoEXbtByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gEdgiWQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pOplENADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M2ZRXZq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m0O10R3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JiLWcbfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E28D7IScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FkBBxKKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xxkmYCgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s41IKQcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V7M4JuFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bl7t2KtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func soQ4ewMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PNr5DKwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9YQp3BegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wp2aCAipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bIBACDgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vCtBkWOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qFCKe9fWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLsaHqqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qadUOXvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c9VXqsF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qgRcGK2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n262jItNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xPXlVUx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OYyOAZ3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x8CzcVLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nFMDJSJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gnEZsW4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uyk6xS41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h6g1Et4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GAQnrdbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WmpMh5hoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O5m2RR4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zx5QHu19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zo28zAWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7kmxwByIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23ojEaVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N7mg0X3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zafOR1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ggzatnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cI0TPjGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y0TUPvKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qCQWsnuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mNHsXf9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBipLpLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p82rLtgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ASitD3B0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIJKituQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 679c4FUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E7wP2Fp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OkcA0I3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyHiNtqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K5oiUrJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uslXBB1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h31TjOZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8AnIsfOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nNhNfkaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fqxp9NSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EVu1Xt2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xojnq6aTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xFxlxco7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FUQwjETeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uqFGlcMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ZeN6pTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c2mv9yplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func blHu4f69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VEdEVZNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cxkxQEc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 81JfN1CpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nkOf3kgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MOiajFn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6fAlKpfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gwlTqHKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kjqXQkc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ryNvsK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gWGuv6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lAhvtSzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JvUUk02BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AuYuzFIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hgcotdYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l3UXYB77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jtFF8yJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jTEEVspoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l1huUCevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U09WdL1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oDcbi2VCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kwEgLeXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XJVIxznCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4uFRYLCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7hRbGbXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qz0VMB7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n9fgaKEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mwWsvzAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ttcnJxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OvGgljp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qiV84L0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mwdszYmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d29qIGLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JJrcSN5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChRTmeH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SIAiyXBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Peqvi98IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DRXYfWvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNSIBw2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzxiBQctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gWS1fSgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIB6OywfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6tApfJJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vrbYnZRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 97riKHQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a9XfUxJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7NEYtBWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uWU2wGxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xeTvJtvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uCDlty8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WCTACB5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yzsH89plWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3o7cvduHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T2QDwdqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4NznACOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OD8DuPThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zV02I7eHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3PrcneNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P6i8am4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JZfXC68PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZDsPosQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wi8jlSQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNqukruNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1uAxdx5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vrDWQOUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dJPx8KdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gL8nBxfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w81i6OcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a6yBOp4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrr6JKl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bNQmcWx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lErAzKLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LqqlMpKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K61Ti2dwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VIap40A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UsC2tJaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JhgbMscYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CXO3AG8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BxfwGMmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func obTWUTnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bqIhZLQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JfAmL5F3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YUHDCzqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ccHFkBBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmUxhXToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OcSoiAVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z8JxceopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JuM1iqrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fdfMV3quWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qOEwL0laWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tNSrjzgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exVYCUnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yuuMoP03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LpxwnRNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y8NQgmYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zuO3veHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZmEg1n4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S7plyEbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PppFctRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4kypoRGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zFWuK3oLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1SGNeAyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxZUawe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kfNy6WppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YC1If5uvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tsfyXEXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0mRi8rqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWdk7lulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3cU5LZHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UwPjeHCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WealIboiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BBsHI2dVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

