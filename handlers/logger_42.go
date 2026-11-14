package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.7" )

func tevkcRk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s561Nm8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBnyCeC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Auca46qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjJ9XMy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5nuGbwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6y3cd7T0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uROzff04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AafQJDQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2nVR6vVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogEcuYBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XoKHqJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NU99HLYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5D8R8hPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7z0jgcdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8RFRhb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7181bxwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpdHs5jbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMoCG4ALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQ79LDaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEs7GOJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIV1T1VQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNMvPvywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoKpd61GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dlo15PiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3noHxY6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAHcQ5GwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjrPXbh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUeRxuQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GwXeB13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBFwcV2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6veT9QzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbJsoep9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRKhZgzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q22gswt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pvFfyLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdKisrsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAG7Vji6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iazkpzWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIuSMzVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCJ6ULQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fusBgb4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGZ3AjxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUrV1HKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71GX1zqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kz6k24ySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huGsoS7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbsntts1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBCgaanxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EY11ekAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kLQVKmgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRopUxEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3D4OJnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FMLn71MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8jO7ri0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99eqAN8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAGtaoocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RnOqsHlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHVYZ1HsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5rcSgtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x31R3c3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXz6xoYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MX5s2o0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jPTRct4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func adu998FDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KacakLp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iZHEeYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKn4WLXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TClHTi3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H463omQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHpQ2p51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wee1aTXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62PCzjT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KVrUQ7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNCfX4unWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5b8bJHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njjJOKT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nDTRMHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vusnh9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQpIQMI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

