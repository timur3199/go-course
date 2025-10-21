package main

import "fmt"

func main(){
	var num int
	fmt.Printf("Введите число: ")
	fmt.Scan((&num))

	if (num >= 12307){
		fmt.Printf("Введденное число %d больше или равно 12307. Результат: %d\n", num, num)
		return
	}

	for num < 12307 {
		if num < 0 {
			num*=-1
		}else if num%7 == 0{
			num*=3
		}else if num%9 == 0{
			num*=13
			num+=1
			continue
		}else{
			num+=2
			num*=3
		}
		
		if (num%9 == 0) && (num%13 == 0){
			fmt.Println("service error")
			return
		}else{
			num+=1
		}
	}

	fmt.Printf("Конечное число: %d\n", num)
	fmt.Printf("Тип результата: %T\n", num)
	
	fmt.Printf("Прописью: %s\n", num_to_word(num))
	fmt.Printf("Тип результата: %T\n", num_to_word(num))
}

func num_to_word(num int) string{
	if num == 0 {return "Ноль"}

	u := []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"}
    t := []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
    h := []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}
    teen := []string{"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}

	th, rem := num/1000, num%1000
    res := ""
    
    if th > 0 {
        res += get_text(th, h, t, teen, u, true) + get_ending(th)
    }
    if rem > 0 {
        if res != "" { res += " " }
        res += get_text(rem, h, t, teen, u, false)
    }
	return res
}

func get_text(num int, h, t, teen, u []string, isTh bool) string {
    if num == 0 { return "" }

parts := []string{}
    if num/100 > 0 { parts = append(parts, h[num/100]) }
    
    rem := num % 100
    if rem == 0 { return join(parts) }
    
    if rem < 10 {
        word := u[rem]
        if isTh && rem == 1 { word = "одна" }
        if isTh && rem == 2 { word = "две" }
        parts = append(parts, word)
    } else if rem < 20 {
        parts = append(parts, teen[rem-10])
    } else {
        parts = append(parts, t[rem/10])
        if rem%10 > 0 {
            word := u[rem%10]
            if isTh && rem%10 == 1 { word = "одна" }
            if isTh && rem%10 == 2 { word = "две" }
            parts = append(parts, word)
        }
    }
    
    return join(parts)

}

func get_ending(num int) string {
    if num%100 >= 11 && num%100 <= 19 { return " тысяч" }
    switch num % 10 {
    case 1: return " тысяча"
    case 2, 3, 4: return " тысячи"
    default: return " тысяч"
    }
}

func join(parts []string) string {
    res := ""
    for i, p := range parts {
        if i > 0 { res += " " }
        res += p
    }
    return res
}