package main

import "fmt"

func SetBit(value int64, position uint, bitValue int) int64 {
    if position > 63 {
        panic("Позиция бита должна быть от 0 до 63")
    }
    
    if bitValue == 1 {
        return value | (1 << position)
    } else if bitValue == 0 {
        return value &^ (1 << position)
    } else {
        panic("Бит может быть только 0 или 1")
    }
}

func GetBit(value int64, position uint) int {
    if (value & (1 << position)) != 0 {
        return 1
    }
    return 0
}

func main() {
    var num int64 = 42 // 101010 в двоичной
    
    fmt.Printf("Исходное число: %d (%b)\n", num, num)
    
    // Устанавливаем 3-й бит в 1 (позиции с 0)
    num = SetBit(num, 3, 1)
    fmt.Printf("После установки 3-го бита в 1: %d (%b)\n", num, num)
    
    // Устанавливаем 1-й бит в 0
    num = SetBit(num, 1, 0)
    fmt.Printf("После установки 1-го бита в 0: %d (%b)\n", num, num)
    
    // Читаем значения битов
    for i := uint(0); i < 8; i++ {
        fmt.Printf("Бит %d: %d\n", i, GetBit(num, i))
    }
}