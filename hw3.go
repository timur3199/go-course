package main

import (
	"errors"
	"fmt"
)

func main(){

	deque := NewDeque()

	fmt.Println("Создаем двухстронюю очередь:")
	fmt.Println(deque)

	deque.PushBack(1).PushFront(2).PushBack(3).PushFront(4).PushBack(-5)
	fmt.Println("После добавления элементов:")
	fmt.Println(deque)

	front, err := deque.Front()
	if err != nil{
		fmt.Println("Ошибка", err)
	}else{
		fmt.Println("Первый элемент:", front)
	}

	back, err := deque.Back()
	if err != nil{
		fmt.Println("Ошибка", err)
	}else{
		fmt.Println("Первый элемент:", back)
	}

	deque.PopFront()
	deque.PopBack()
	fmt.Println("После удаления элементов:")
	fmt.Println(deque)

	deque.PushBack(100).PushFront(200).PushBack(300).PushFront(400).PushBack(-500)
	fmt.Println("После добавления новых элементов:")
	fmt.Println(deque)
	
	fmt.Println("Обход с помощью итератора:")
	iterator := deque.Iterator()
	for{
		value, hasMore := iterator()
		if !hasMore{
			break
		}
		fmt.Printf("%v ", value)
	}
	fmt.Println()

	fmt.Printf("Размер очереди: %d\n", deque.Size())
	fmt.Printf("Очередь пуста: %t\n", deque.IsEmpty())

	deque.Clear()
	fmt.Println("После очистки:")
	fmt.Println(deque)
	fmt.Printf("Размер очереди: %d\n", deque.Size())
	fmt.Printf("Очередь пуста: %t\n", deque.IsEmpty())
}

type Node struct{
	data interface{}
	prev *Node
	next *Node
}

type Deque struct{
	head *Node
	tail *Node
	size int
}

func NewDeque() *Deque{
	return &Deque{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (d *Deque) PushFront(value interface{}) *Deque{
	newNode := &Node{
		data: value,
		prev: nil,
		next: d.head,
	}

	if d.IsEmpty(){
		d.head = newNode
		d.tail = newNode
	}else{
		d.head.prev = newNode
		d.head = newNode
	}
	d.size++
	return d
}

func (d *Deque) PushBack(value interface{}) *Deque{
	newNode := &Node{
		data: value,
		prev: d.tail,
		next: nil,
	}

	if d.IsEmpty(){
		d.head = newNode
		d.tail = newNode
	}else{
		d.tail.next = newNode
		d.tail = newNode
	}
	d.size++
	return d
}

func (d *Deque) Front() (interface{}, error){
	if d.IsEmpty(){
		return nil, errors.New("Front: попытка прочитать элемент из пустой очереди")
	}
	return d.head.data, nil
}

func (d *Deque) Back() (interface{}, error){
	if d.IsEmpty(){
		return nil, errors.New("Back: попытка прочитать элемент из пустой очереди")
	}
	return d.tail.data, nil
}

func (d *Deque) PopFront() (interface{}, error){
	if d.IsEmpty(){
		return nil, errors.New("PopFront: попытка удалить элемент из пустой очереди")
	}
	
	value := d.head.data
	if d.size == 1{
		d.head = nil
		d.tail = nil
	}else{
		d.head = d.head.next
		d.head.prev = nil
	}
	d.size--
	return value, nil
}

func (d *Deque) PopBack() (interface{}, error){
	if d.IsEmpty(){
		return nil, errors.New("PopBack: попытка удалить элемент из пустой очереди")
	}
	
	value := d.tail.data
	if d.size == 1{
		d.head = nil
		d.tail = nil
	}else{
		d.tail = d.tail.prev
		d.tail.next = nil
	}
	d.size--
	return value, nil
}

func (d *Deque) IsEmpty() bool{
	return d.size == 0
}

func (d *Deque) Clear(){
	for d.size > 0{
		d.PopFront()
	}
}

func (d *Deque) Size() int{
	return d.size
}

//строковое представление очереди(нужно для форматированного вывода очереди и для удобства отладки)
func (d *Deque) String() string{
	if d.IsEmpty(){
		return "[empty]"
	}
	res := "[ "
	curr :=d.head
	for curr != nil{
		res += fmt.Sprintf("%v ", curr.data)
		curr = curr.next
	}
	res += "]"
	return res
}

//итератор для обхода нашей очереди
func (d *Deque) Iterator() func() (interface{}, bool) {
	curr := d.head
	return func() (interface{}, bool) {
		if curr == nil {
			return nil, false
		}
		value := curr.data
		curr = curr.next
		return value, true
	}
}
