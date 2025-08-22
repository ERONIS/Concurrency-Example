package main

import (
	"context"
	miner "cunncurency/Miner"
	postmam "cunncurency/Postman"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var coal atomic.Int64
	var mails []string

	mailsContext, mailsCancel := context.WithCancel(context.Background())
	coalContext, coalCancel := context.WithCancel(context.Background())

	coalTransferPoint := miner.MinerPool(coalContext, 2)
	mailTransferPoint := postmam.PostmanPool(mailsContext, 2)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("-------->>>>  Мой рабочий день шахтера закончен")
		mailsCancel()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("-------->>>>  Мой рабочий день почтальона закончен")
		coalCancel()

	}()

	wg := &sync.WaitGroup{}
	mtx := sync.Mutex{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()
	wg.Wait()

	fmt.Println("Сегодня было добыта ", coal.Load(), "угля")
	mtx.Lock()
	fmt.Println("Сегодня было доставлено", len(mails), "письем")
	mtx.Unlock()

}
