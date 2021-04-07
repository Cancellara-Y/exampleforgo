package queue

import (
	"sync"
)

type QueueV1 struct {
	q *Queueimpl7
	sync.Mutex
}

func NewSafe() *QueueV1 {
	sq := &QueueV1{
		q: New(),
	}

	return sq
}

func (s *QueueV1) Len() int {
	s.Lock()
	n := s.q.Len()
	s.Unlock()
	return n
}

func (s *QueueV1) Push(v interface{}) {
	s.Lock()
	defer s.Unlock()

	s.q.Push(v)
}

func (s *QueueV1) Pop() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	return s.q.Pop()
}

func (s *QueueV1) Front() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	return s.q.Front()
}