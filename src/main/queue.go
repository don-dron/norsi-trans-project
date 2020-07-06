package main

import (
	"sync/atomic"
	"unsafe"
)

type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	len  uint64
}
type directItem struct {
	next unsafe.Pointer
	v    interface{}
}

func NewQueue() *Queue {
	head := directItem{next: nil, v: nil}
	return &Queue{
		tail: unsafe.Pointer(&head),
		head: unsafe.Pointer(&head),
	}
}

func (q *Queue) Enqueue(v interface{}) {
	i := &directItem{next: nil, v: v}
	var last, lastnext *directItem
	for {
		last = loaditem(&q.tail)
		lastnext = loaditem(&last.next)
		if loaditem(&q.tail) == last {
			if lastnext == nil {
				if casitem(&last.next, lastnext, i) {
					casitem(&q.tail, last, i)
					atomic.AddUint64(&q.len, 1)
					return
				}
				casitem(&q.tail, last, lastnext)
			}
		}
	}
}

func (q *Queue) Dequeue() interface{} {
	var first, last, firstnext *directItem
	for {
		first = loaditem(&q.head)
		last = loaditem(&q.tail)
		firstnext = loaditem(&first.next)
		if first == loaditem(&q.head) {
			if first == last {
				if firstnext == nil {
					return nil
				}
				casitem(&q.tail, last, firstnext)
			} else {
				v := firstnext.v
				if casitem(&q.head, first, firstnext) {
					atomic.AddUint64(&q.len, ^uint64(0))
					return v
				}
			}
		}
	}
}

func (q *Queue) Length() uint64 {
	return atomic.LoadUint64(&q.len)
}
func loaditem(p *unsafe.Pointer) *directItem {
	return (*directItem)(atomic.LoadPointer(p))
}
func casitem(p *unsafe.Pointer, old, new *directItem) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
