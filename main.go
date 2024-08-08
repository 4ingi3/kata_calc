package main 

import "fmt"; "strconv" 

func romanToArabic(roman string) int {   // преобразуем римские числа в арабские 
  romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
  result := 0
  prevValue := 0

  for i := len(roman) - 1; i >= 0; i-- {
    value := romanNumerals[rune(roman[i])]
    if value == 0 {
        return -1
    }
    if value < prevValue {
        result += value
    }
    prevValue = value 
    }
  return result
  }
  if result < 1 || result > 10 {
        return 0, fmt.Errorf("число вне допустимого диапазона 1-10")
  }
    return result, nil
}
    
  
    func arabicToRoman(arabic int) string {   // преобразуем арабские числа в римские 
  arabicNumerals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
  romanNumerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
  result := ""

  for i, arabicValue := range arabicNumerals {
    for arabic >= arabicValue {
        arabic -= arabicValue 
        result += romanNumerals[i]
      }
    }
  return result 
  }

func main() {
    var input1, input2, operation string 

  fmt.Print("Введите выражение (например, 2 + 2 или II + III): ")
  fmt.Scanf("%s %s %s", &input1, &operation, &input2)

  isRoman := func(s string) bool {  // функция проверки на римские числа
    for _, char := range s {
      if _, ok := map[rune]int{'I': 1, 'V': 5, 'X': 10}[char]; !ok {
        return false
      }
    }
    return true
  }

  if isRoman(input1) && isRoman(input2) { // проверяем являются оба значения римскими числами
    a := romanToArabic(input1)
    b := romanToArabic(input2) 

    if a == -1 || b == -1 {
      fmt.Println("Ошибка: Некорректное римское число.")
      return
    }

      processOperation(a, b, operation, true)
   } else }   // обрабатываем арабские числа 
      a, err1 := strconv.Atoi(input1)
      b, err2 := strconv.Atoi(input2)

      if err1 != nil || err2 != nil || a < 1 || a > 10 || b < 1 || b > 10 { 
        fmt.Println("Ошибка: Неккоректное арабское число.")
        return
      }

      processOperation(a, b, operation, false) 
   }
}

  func processOperation(a, b int, operation string, isRoman bool) { // функция для обработки операции
    var result int 
    switch opearion {
      case "+":
          result = a + b 
      case "-":
          result = a - b 
      case "*":
          result = a * b 
      case "/":
          if b == 0 {
            fmt.Println("На ноль делить нельзя!")
            return
          }
          result = a/b
      default: 
            fmt.Println("Неверная операция. Доступные операции: +, -, *, /.")
            return
      }

      if isRoman {
        fmt.Println(arabicToRoman(result))
      } else }
        f,t.Println(result)
      }
}

            


                           
                           
