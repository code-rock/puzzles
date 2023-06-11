// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

func doSomething(ctx context.Context) error {
	// Бесконечный цикл
	for {
		// Слушает если контекст выполнения вдруг остановиться по какой-то причине
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		log.Println("do some work")
		select {
		case <-ctx.Done():
			// Если контекст истекает, выбирается этот случай
			return errors.New("interrupted in the middle of the work")

		case <-time.After(time.Second):
			// Когда работа завершается до отмены контекста
			log.Println("done")
		}

	}
}

func main() {
	// Передаем контекст с тайм-аутом, чтобы сообщить блокирующей функции, что он
	// должен прекратить свою работу по истечении времени ожидания.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Возвращает копию родительского контекста, который помечен как выполненный
	// (его канал Done закрыт), когда поступает один из перечисленных сигналов,
	// когда вызывается возвращаемая функция остановки или когда закрывается канал
	// Done родительского контекста, в зависимости от того, что произойдет раньше.
	ctx, cancel = signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	// Обеспечивает синхронизацию, распространение ошибок и отмену контекста для
	// групп горутин, работающих над подзадачами общей задачи.
	group, ctx := errgroup.WithContext(ctx)

	log.Println("starting")
	// Запуск группы горутин которая что-то делает
	group.Go(func() error {
		return doSomething(ctx)
	})

	// Ожидание выполнения и если произойдет ошибка вывод ее в лог
	if err := group.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("stopped")
}
