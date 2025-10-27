package main

import(
	"errors"
	"fmt"
)

func main(){

	stack := NewStack()

	fmt.Println("Создаем стек:")
	fmt.Println(stack)

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push("hi")
	stack.Push(3.14)
	fmt.Println("После добавления элементов:")
	fmt.Println(stack)

	top, err := stack.Peek()
	if err !=nil{
		fmt.Println("Ошибка:", err)
	}else{
		fmt.Println("Вершина стека:", top)
	}

	fmt.Println("\nИзвлекаем элементы:")
	for !stack.IsEmpty(){
		value, err := stack.Pop()
		if err != nil{
			fmt.Println("Ошибка:", err)
		}else{
			fmt.Printf("Извлечено: %v\n", value)
			fmt.Printf("Текущий размер: %d\n", stack.Size())
		}
	}

	_,err = stack.Pop()
	if err != nil{
		fmt.Println("\nОшибка при извлечении:", err)
	}

	fmt.Println("\nДобавляем элементы снова...")
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)
	fmt.Println("До очистки:", stack)

	stack.Clear()
	fmt.Println("После очистки:", stack)
	fmt.Printf("Стек пуст: %t\n", stack.IsEmpty())
}

type Stack struct{
	elements []interface{}
}

func NewStack() *Stack{
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

func (s *Stack) Peek() (interface{}, error){
	if s.IsEmpty(){
		return nil, errors.New("Peek: попытка прочитать элемент из пустого стека")
	}
	return s.elements[len(s.elements) - 1], nil
}

func (s *Stack) Push(value interface{}){
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (interface{}, error){
	if s.IsEmpty(){
		return nil, errors.New("Pop: попытка извлечь элемент из пустого стека")
	}

	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value, nil
}

func (s *Stack) IsEmpty() bool{
	return len(s.elements) == 1
}

func (s *Stack) Size() int{
	return  len(s.elements)
}

func (s *Stack) Clear(){
	s.elements = make([]interface{}, 0)
}

//строковое представление стека
func (s *Stack) String() string {
	if s.IsEmpty() {
		return "Stack: [empty]"
	}
	return fmt.Sprintf("Stack: %v", s.elements)
}