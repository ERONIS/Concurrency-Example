package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Работа оконченв")
			return
		default:
			fmt.Println("Я шахтер номер ", n, "начал добывать уголь")
			time.Sleep(1 * time.Second)
			fmt.Println("Я шахтер номер ", n, "Добыл уголь", power)
			transferPoint <- power
			fmt.Println("Я шахтер номер ", n, "Передал уголь", power)
		}

	}

}

func MinerPool(ctx context.Context, countMiner int) <-chan int {
	coalTransferPoint := make(chan int)
	wg := &sync.WaitGroup{}

	for i := 0; i < countMiner; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}
	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint

}
