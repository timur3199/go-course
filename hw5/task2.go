package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main(){

	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	_, err := fmt.Scan(&numWorkers)
	if err != nil || numWorkers <= 0{
		log.Fatal("Некоректное количество воркеров")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++{
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func ()  {
		counter := 1
		for{
			select{
			case <-ctx.Done():
				close(jobs)
				return
			default:
				jobs <- counter
				counter++
				time.Sleep(50*time.Millisecond)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nПолучен сигнал завершения, завершаю работу...")

	cancel()
	wg.Wait()

	fmt.Println("Все воркеры завершили работу. Программа завершена.")
}

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup){
	defer wg.Done()

	for{
		select{
		case job, ok := <-jobs:
			if !ok{
				fmt.Printf("Воркер %d: канал закрыт, завершаю работу\n", id)
				return
			}
			fmt.Printf("Воркер %d: обработал задание %d\n", id, job)
			time.Sleep(100*time.Millisecond)
		case <-ctx.Done():
			fmt.Printf("Воркер %d: получил сигнал завершения\n", id)
			return
		}
	}
}