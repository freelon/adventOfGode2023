package util

type Queue[T any] struct {
	messages []T
}

func (q *Queue[T]) Enqueue(m T) {
	q.messages = append(q.messages, m)
}

func (q *Queue[T]) Dequeue() (message T, ok bool) {
	if len(q.messages) > 0 {
		m := q.messages[0]
		q.messages = q.messages[1:]
		return m, true
	} else {
		var x T
		return x, false
	}
}
