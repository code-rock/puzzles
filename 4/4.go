// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные
// из канала и выводят в stdout. Необходима возможность выбора
// количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func runConsumer(dataCh <-chan interface{}) {
	// Читает из канала и пишет в stdout
	for value := range dataCh {
		io.WriteString(os.Stdout, fmt.Sprintf("%d", value))
	}
}

func runProducer(dataCh chan<- interface{}, stopCh <-chan struct{}) {
	data := uint64(0) // TODO: generate something more interesting
	for {
		// Попытка остановить максимально быстро
		select {
		case <-stopCh:
			return
		default:
		}

		select {
		// Дублирование на случай если оба процесса окажутся активными
		case <-stopCh:
			return
		// Запись информации в канал
		case dataCh <- data:
			// Увеличение числа
			data++
		}
	}
}

func runService(workersCount int, stopCh <-chan struct{}) {
	// Переменная группы ожидания
	var wg sync.WaitGroup
	// Создание канала с заданием количества воркиров
	dataCh := make(chan interface{}, workersCount) // the buffer is optional
	// Перебираем возможные ключи воркиров
	for i := 0; i < workersCount; i++ {
		// Добовляем горутину в счетчик ожидания группы
		wg.Add(1)
		// Запуск горутины
		go func() {
			// Даст понять группе, что горутина закончина после выполнения функции
			defer wg.Done()
			// запуск функции слушетеля информации из канала
			runConsumer(dataCh)
		}()
	}
	// запуск функции добовления информации в канал
	runProducer(dataCh, stopCh)
	// закрытие канала
	close(dataCh)
	// Ждет окончания горутин
	wg.Wait()
}

// Функция завершения сервиса
func createStopChannel() <-chan struct{} {
	// Создание канала
	stopCh := make(chan struct{})
	// создание канала для поулчения сигналов системы
	sigintCh := make(chan os.Signal, 1)
	// регистрирует данный канал для получения уведомлений о сигнале
	signal.Notify(sigintCh, syscall.SIGINT)

	go func() {
		// Закрываем каналы и отписываемся от отслеживания ctrl + c
		// после вып функции в обратном порядке
		defer close(stopCh)
		defer close(sigintCh)
		defer signal.Reset(syscall.SIGINT)
		// Получаем знак нажатия комбинации клавиш
		<-sigintCh
	}()
	// Возаращение канала остановки
	return stopCh
}

func main() {
	// Переменная для хранения количества воркеров
	var workersCount int
	// Привязка переменной к атрибута командной строки
	flag.IntVar(&workersCount, "w", 1, "the number of workers")
	// Разбора командной строки на определенные флаги
	flag.Parse()
	// Вывод ошибки при открицательном задании количества воркиров
	if workersCount <= 0 {
		log.Fatalf("cannot start the program with %d workers", workersCount)
	}
	// Запись в переменную канад для остановки
	stopCh := createStopChannel()
	// Передаем в функцию запуска сервиса количество воркиров и канал для
	// отслеживания остановки программы
	runService(workersCount, stopCh)
}
