package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func newList() *list {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.first
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if l.len == 0 {
		l.first = item
		l.last = item
		item.Prev = nil
		item.Next = nil
		l.len++
		return item
	}
	temp := l.first
	l.first.Prev = item
	l.first = item
	l.first.Next = temp
	l.first.Prev = nil
	l.len++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := new(ListItem)
	item.Value = v
	if l.len == 0 {
		l.first = item
		l.last = item
		item.Prev = nil
		item.Next = nil
		l.len++
		return item
	}
	l.last.Next = item
	item.Next = nil
	item.Prev = l.last
	l.last = item
	l.len++
	return item
}

func (l *list) Remove(i *ListItem) {
	if l.len == 0 {
		return
	}
	if l.len == 1 {
		l.first = i
		l.last = i
		i.Value = nil
		l.len--
		return
	}
	if i == l.first {
		l.first = i.Next
		i.Next = nil
		l.first.Prev = nil
		l.len--
		return
	}
	if i == l.last {
		l.last = i.Prev
		l.last.Next = nil
		i.Prev = nil
		l.len--
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	l.len--
	i.Prev = nil
	i.Next = nil
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.first {
		return
	}
	if i == l.last {
		l.last = l.last.Prev
		l.last.Next = nil
		temp := l.first
		l.first = i
		l.first.Prev = nil
		l.first.Next = temp
		temp.Prev = l.first
		return
	}
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	i.Prev = nil
	i.Next = l.first
	l.first = i
}
