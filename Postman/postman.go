package postmam

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(ctx context.Context, wg *sync.WaitGroup,
	transferPoint chan string, n int, mail string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон номер ", n, "Мой рабочий день закончен")
			return
		default:
			fmt.Println("Я почтальон номер  ", n, "Я взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я почтальон номер  ", n, "Я донес письмо")
			transferPoint <- mail
			fmt.Println("Я почтальон номер  ", n, "Я отдал письмо")
		}
	}

}

func PostmanPool(ctx context.Context, PostmanCount int) chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 0; i <= PostmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}
	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()
	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	pts := map[int]string{
		1: "Привет от друга",
		2: "Приглашение на свадьбу",
		3: "Скида на маркетплейс",
	}
	mail, ok := pts[postmanNumber]
	if !ok {
		return "Вы выиграли лотерею"
	}
	return mail

}
