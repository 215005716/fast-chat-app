package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.5" )

func 0HKaj8QQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YI6jDUvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zy0oWasnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqXoWsVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zPu6UWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7jlTd4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func skjnc7ISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXUbr2wZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXPHZ7ltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mc0tuvbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACOzdUo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20S4pYMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7E0FjkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPVwlf4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93d7AoukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0I4UrwN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKTsVRz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUr5kPG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHNpS2uFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fp9JOmd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USz4nKxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V02tXYV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0W3nLhoUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upCZqOdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2WwT0JnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kzsv4WeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLO3fep6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Te0wmGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcI2jSozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HN3FJJO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3kniawWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2OGVdvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySoyVpSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmCQLvNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEOqMvIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbWKstkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJ5D9qFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtkXP18yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9C7MYYeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EE7PA9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hTOHwI5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uqZGXeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIxqqgbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrSCuOK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPoU6HSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejcRYSXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nbKgtayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGIB5GhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OFkW7o5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEVJ9rb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVSAguEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZspDN9myWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pNa0I3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWo9MuZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgAMshSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54Na4S6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCyrL4H8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iz5l6DZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dx0WQaQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iws96IZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0hagJ9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjyI6E8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwJ2broWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLzEY9bKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GysujRueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItaTmqLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJvy7Mh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVCEU2sGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxP5GFhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2s9LgXuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlTDKIYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhX1JxmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rWwpuSUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xC2Z2wrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zYu6SepTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9D8wxpF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5sfgy3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdyeOLbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ne14JX2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aYJZuCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdDcGMlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wT6CPQrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOBebykPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rz1SWBfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yk1cXIh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xn0qEfxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkSKgVwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wY5YtWbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xtUvw0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IngUf8KGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCpH3VFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFuZgD8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1f1HJhRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Osp7PQpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vddqqpN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfIDXKfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzxbMS7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBvKUgdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQ6QWTLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5Pl9BifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FhmmrmY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3gjc4r9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBHnQ9EIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Li8nZ9OYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1mPa7gCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDF8ZRFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5x14rEyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JONF66GBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zUVrnDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dA4jMQGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L12B0u2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ied0wVL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUEFHov1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dLq5OaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yS0xj9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WtaPauxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d2BM4kKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EM8caiLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2HlsnMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvOFK8P6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nr5Gv5mXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRbaa57gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiZ6GoIZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func quasgbGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06EuLEq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AB5uKlQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9M8VQ9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjLHNm0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIxfdp1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6MagE8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqP4VsK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iKtMAk9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weDKoi1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTDto9HIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGY15LCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwJVdELHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmYlu683Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rU5xdyHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iDqpu60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VblicNCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iCAfGrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AqhaYVFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZ4tjRhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fz4ZtCzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfFq2407Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SrJTx8e9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBSZjZOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0YN1e5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFfh9rmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 707X133rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZpbi9mHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8PjqesRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUpxCHrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDBQFdDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnzX7P0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5ApLULWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9Un5rL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjibgKxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJ9oi8tzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSd1m6auWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ii9G3LjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zhp3LPpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qghEj3fVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rBSjHL46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YWPzcXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4xykLvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rv7RB6PzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Xj519OIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QG6dcqWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDsteKkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Im5sNXAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKRU5wPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oJtZjNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVeRWR8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qayziw9TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 618iwcQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbZQVpW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgTi69QeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdU5Qo8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JgKvlzSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkKe89CgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0diGWqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCBpcEjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXmHDNX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5HxkTacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmJ0ZcU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4VDxXUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwCOMwUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5xqDTPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xC7sbf9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23jGwsdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHsq8cgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73xO3Uk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDUMPYMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hb9WARX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJxMSWUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBNlJf34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8NA4m9ZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjF3Eet5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2fBQ1YsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQrMmmzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPF4gNKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5B07l5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIgJGU46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zaaix4orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fzk5rnpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stcMTNRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lR4K8IzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjSke7TYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nkr1JtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tJOwU5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XkOqSTBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bxdK1DNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOOzVv3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qw9ojlj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7AbCqvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKGFVmtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vENIdFpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2SbOFTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ByFDh7NLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VioKwofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woNtbUKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VBza6orWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wMg97YCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iV4nAP2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvRsuixdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7r5iD9O7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHoe9TXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I857EvJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ipBKM7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlUnHFv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwSFwdfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4s0taUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PR8sh7GEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfPOBPOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyBjgdLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1qHiWeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnOvPEXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CL1KydbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECY2QNJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKJ2CaZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

