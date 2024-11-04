package utils

import "strings"

type MagicNumber struct {
    Number int
}

func (m *MagicNumber) Multiply(n int) {
    m.Number *= n
}

func MagicSum(n int) int {
    return n + n
}

func MagicPow(n int) int {
    result := 1
    for i := 0; i < n; i++ {
        result *= n
    }
    return result
}

func MagicOdd(n int) bool {
    return n%2 != 0
}

func MagicGrade(n int) string {
    grades := map[int]string{
        0: "Zero",
        1: "Bad",
        2: "Ok",
        3: "Nice",
        4: "Awesome",
        5: "Exceptional",
    }
    if grade, exists := grades[n]; exists {
        return grade
    }
    return "Unknown"
}

func MagicName(n int) []string {
    name := "radit"  
    return strings.Split(strings.Repeat(name+",", n), ",")[:n]
}

func MagicTria(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

func MagicChange(n *int) {
    *n *= 2
}
