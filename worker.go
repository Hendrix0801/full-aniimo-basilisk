package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.6" )

func UfbJb5TDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7anpcNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2ZhpRO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVGI0SXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxAZWC3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMhtMPXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRvY6iQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXJkRIjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vzOdyNT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vD8iSq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwrlDt3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzU2u4E1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJiyigNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69p3hwWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EEOMWPS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCONvqLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJ9eXfvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4tVTBI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imRJwN1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmIYuOvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdoCLOmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePBt0wmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JxhU90jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6gER6i0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cz7BEZ9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miA0xsAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LG5b5pw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15UaOnplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peNWZt5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kzq5PvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uNtZkyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9ro8jFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTxtvXADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K3F4K6sEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKESAZcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sq2NbI2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgzJjpCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yobx9rRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlTItYp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uT1yAtHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1MhXoBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXEWZ2QKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4ZwfbpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3bB7IIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UY0YRRzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEtqJh53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2k6W37rIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1f616EGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0LoWeSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H15Fhc0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhYxAQ6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dY3Shr1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AowccblnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pJUo012Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p75HmEn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kbrBAeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUJPrMLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPdYMMBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3733VCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cntzVPIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyVI9hQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KudXV9w8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqG4d7dJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVMFB4XBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62t7nNXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RcVADs1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SzucYGt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hchdcqfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMuWVvdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhl8zI2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2Y0MhobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JryPJIukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKL3stnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQtZQygSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDTGT4y7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func anXf7RdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Mubs2URWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4VtJH8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u01gNC0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVNKbI1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyKhikpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQYGtw69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yaIjJPVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDgFeOTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hansTd0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TW6oeZfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8TsQz7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5kpjYmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99QDEXDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUTsf1KFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8h8FiMCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dea9Vlx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ar3M14kbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rji5QENdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HoG39icnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdvs2hmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXJJsa2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nDAHDbbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLQOljC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inUKlQGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOssO7OyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZVrWiU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsi9gwwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpN99esgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8c089qdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QH5rLWs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K65045meWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PsdVVsmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvetXGx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVIr8oc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pf1Bm5srWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptIRWYf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I29dtF4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Rb7JIKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOov2wE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJrMyOQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uGnwMMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duf4hCjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpLAaWq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbxaC9gsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U2PATOGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VoKG5dG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTURamQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OGfqtot2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkRYhMteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOU9AlFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AUBKR5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkfOJb6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swDf4yVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmvbwpwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func douGkfTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9GkozVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aRAfuGQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdb0iFBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPuYxGJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HZqEBecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snalwdRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OnhmQkACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MkdwZw0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VO1aXujoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRW2kJyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3gnZetuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82K4W9wnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rszf9T2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSDjaiBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func at7lddiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9FDijsmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHPUqmggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGbgHl7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50FtLhxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJ4KbhOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVeitnArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t40fyR8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lT4XGbkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func quYoKndIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7Q0PdkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1GMElirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmLkU16pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Umln7Vd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3hwvtbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwwLlhYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHqpHbRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nXn15GzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7eCjvJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2vuceEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BViB41jeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQdwHS2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owyaHI2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNiHj7FcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHoRU2LJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Pe0OPGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7aTNuHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eDEkJBLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWFTDoNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyWvSW8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bm6ff9Q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCSt3kccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhqC5M71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CgojYa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j2UOQ62aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oF66zfTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHvHfV4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBGpPs4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJUtzYhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32KtRVvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdkLaCPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2iqQrnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P35yfFIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOT8TpttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvTdALJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zg9BvBSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GcHV8AtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SpIL7twWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90vCI9RQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fsif8kWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoGGjnm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kc3mY33TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hQUeMyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsZeHbRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NogwcEiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VdmT0b8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W45yFELiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5qvRCOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BIbzqEr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uWpBDSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrVsIX3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TtNlOFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjQOofRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rw34CSbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eaCvYXGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfj0D6I8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mfmAiRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ih7ucCvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBAFGHjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkCuLhMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQWqAOCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B381H7PjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhBtyJ06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkmuk8ZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtebRk47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RF413W2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func azuQVoTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VszX0KFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78oveuILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1G4BmUTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZdMjrNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7PuxW1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyr9mmNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxJV6Z34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvKlT8JfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pj5fpvKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdzwbL76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPHQnBvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtL4eykJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A9UM8p33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEnu3JwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fkODpT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2DCixFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZrKkWTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpfSUNmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dudkNGncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func echXdcsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2033yIhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4r6MENBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kdb4rayDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqXVcHixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h47J5FbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5f6hb5qfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrzE1zs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUMcS3tvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNalo7ECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wfJlIyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HD2aCEwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nluzxCO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFtDC4n7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juXM02Z5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DK1aA51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32AXvlKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SE0SetZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sH8CbwDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uwqPLK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDzT4Oq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jAustNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3w2GjbEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjufQXmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZFWmv37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wcBNtWVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZn2rmqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypz1t9QwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVLIQsWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OmOqX6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXX7TugNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGwE3nxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zekLC5d2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ergq8l5FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trVWoSXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 07EADtmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func krXWwFCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLgtnCccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHcz1cSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Epn5FEFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lRhtFRgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A411XHtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayGFS3m0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBIqoy4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPQKtzkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZg7YIWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0V0krAMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkyZZiQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXcX8C7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVEtaEtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkHlsS8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHA1dDFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLSnvZGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57aJOIaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwMWEEpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvKL3QVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KM0ZjOpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jWxjGcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5f9gGcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtTr497EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtbbKYxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZ5UgZDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2cgXKKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGqr1MREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMKkc7bkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxibVbdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EInASPDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8x7ji1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjkRXOfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bf7HGTzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COUAp2ERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dc4M3LCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X36E7KtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2xV3dcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZSGHgrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSdwM5X8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

