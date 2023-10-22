// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

package main
var justString string

func createHugeString(size int) string {
    // Создаем большую строку
    for i 
    hugeString := `On architectures lacking a dedicated swap instruction, because it avoids the extra temporary register, the XOR swap algorithm is required for optimal register allocation. 
    This is particularly important for compilers using static single assignment form for register allocation; these compilers occasionally produce programs that need to swap two registers when no registers are free. 
    The XOR swap algorithm avoids the need to reserve an extra register or to spill any registers to main memory.[5] The addition/subtraction variant can also be used for the same purpose.[6]
    This method of register allocation is particularly relevant to GPU shader compilers. On modern GPU architectures, spilling variables is expensive due to limited memory bandwidth and high memory latency, 
    while limiting register usage can improve performance due to dynamic partitioning of the register file. The XOR swap algorithm is therefore required by some GPU compilers` // Здесь создайте большую строку

    // Возвращаем только первые 100 символов
    return hugeString[:100]
}

func someFunc() {
    v := createHugeString(1 << 10)
    justString = v
}

func main() {
    someFunc()
}