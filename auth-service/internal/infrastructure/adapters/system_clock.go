package adapters

import "time"

// TODO: Адаптер времени
//  - Получать текущее время
//  - Нужен для мокирования времени в тестах

type SystemClock struct{}

func (SystemClock) Now() time.Time {
	return time.Now()
}
