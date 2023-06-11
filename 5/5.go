// Разработать программу, которая будет последовательно отправлять
// значения в канал, а с другой стороны канала — читать. По истечению
// N секунд программа должна завершаться.

package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

func runConsumer(dataCh <-chan interface{}) {
	// Читает из канала последовательно
	for data := range dataCh {
		// Пишет в лог
		log.Printf("received data '%v'", data)
	}
}

func runProducer(dataCh chan<- interface{}, stopCh <-chan struct{}) {
	// Создаие переменной которая цифру хранит положительную
	var data uint64
	// Бесконечный цикл
	for {
		select {
		// Слушает данные из канала остановки, когда время выйдет
		case <-stopCh:
			return
		default:
		}
		// Пишет в консоль что посылает цифру
		log.Printf("sending data '%v'", data)
		select {
		// Слушает данные из канала остановки, когда время выйдет
		case <-stopCh:
			return
		// Пишет цифру в канал
		case dataCh <- data:
			// Делет цифру больше на единичку на следующем шаге
			data++
		}
	}
}

func createStopChannel(timeout time.Duration) <-chan struct{} {
	// Создание канала остановки
	stopCh := make(chan struct{})
	// Создает новый таймер, который будет отправит текущее время на свой канал,
	// по окончанию времени задержки
	timer := time.NewTimer(timeout)
	// Запуск горутины
	go func() {
		// Закрытие конала по выходу из функции
		defer close(stopCh)
		// Остановка таймера по выходу из функции
		defer timer.Stop()
		// Чтение времени конца таймера из канала
		<-timer.C
	}()

	// Возврат канала из функции
	return stopCh
}

func main() {
	var timeout uint
	flag.UintVar(&timeout, "t", 1, "the timeout in seconds")
	flag.Parse()

	log.Println("starting the service")
	// Создание группы ожидания горутин
	var wg sync.WaitGroup
	// Создание канала с данными в любом формате
	dataCh := make(chan interface{})
	// Добовляем в счетчик горутин одну
	wg.Add(1)
	// Запускаем горутину
	go func() {
		// Говорим что горутина исполнена в самом конце перед выходом из функии группе горутин
		defer wg.Done()
		// Запускаем потребителя и передаем канал
		runConsumer(dataCh)
	}()
	// Создание функции для остановки программы задание в ней таймера
	// можно использовать Context.WithTimeout вместо
	stopCh := createStopChannel(time.Duration(timeout) * time.Second)
	// Запуск поставщика данных в канал, принимает канал для данных и канал остановки
	runProducer(dataCh, stopCh)
	// Закрытие канала с данными
	close(dataCh)
	// Ожидание выполнени горутин
	wg.Wait()
	log.Println("the service stopped")
}
