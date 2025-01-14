package tasks

import (
	"strconv"
	"sync"
)

type List struct {
	l  sync.Mutex
	m  map[string]string
	id int
}

func NewList() *List {
	return &List{
		m: map[string]string{},
		l: sync.Mutex{},
	}
}

func (l *List) Add(task string) string {
	l.l.Lock()
	defer l.l.Unlock()
	l.id = l.id + 1
	newId := strconv.Itoa(l.id)
	l.m[newId] = task

	return newId
}

func (l *List) Del(id string) bool {
	l.l.Lock()
	defer l.l.Unlock()
	_, ok := l.m[id]

	if !ok {
		return false
	}

	delete(l.m, id)

	return true
}

func (l *List) All() map[string]string {
	l.l.Lock()
	defer l.l.Unlock()

	all := make(map[string]string, len(l.m))

	for k, v := range l.m {
		all[k] = v
	}

	return all
}
