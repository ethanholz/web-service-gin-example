package lib

import "errors"

type Node struct {
	Album Album
	next  *Node
	prev  *Node
}

type List struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (L *List) RemoveHead() (*Album, error) {
	if L.Head == nil && L.Tail == nil && L.Length == 0 {
		return nil, errors.New("List is empty")
	}
	L.Length--
	frontNode := L.Head
	if L.Head.prev == nil && L.Head.next == nil {
		L.Tail = nil
		L.Head = nil
		return &frontNode.Album, nil
	}
	newHead := frontNode.next
	newHead.prev = nil
	L.Head = newHead
	return &frontNode.Album, nil
}

func (L *List) RemoveTail() (*Album, error) {
	if L.Head == nil && L.Tail == nil && L.Length == 0 {
		return nil, errors.New("List is empty")
	}
	L.Length--
	lastNode := L.Tail
	if L.Tail.prev == nil && L.Tail.next == nil {
		L.Tail = nil
		L.Head = nil
		return &lastNode.Album, nil
	}
	newTail := lastNode.prev
	newTail.next = nil
	L.Tail = newTail
	return &lastNode.Album, nil
}

func (L *List) InsertStart(album Album) {
	newNode := &Node{
		Album: album,
		next:  nil,
		prev:  nil,
	}
	L.Length++
	if L.Tail == nil && L.Head == nil {
		L.Head = newNode
		L.Tail = newNode
		return
	}
	firstNode := L.Head
	firstNode.prev = newNode
	newNode.next = firstNode
	L.Head = newNode
}

func (L *List) Insert(album Album) {
	newNode := &Node{
		Album: album,
		next:  nil,
		prev:  nil,
	}
	L.Length++
	if L.Tail == nil && L.Head == nil {
		L.Tail = newNode
		L.Head = newNode
		return
	}
	lastNode := L.Tail
	lastNode.next = newNode
	newNode.prev = lastNode
	L.Tail = newNode
}

func (L *List) GetNodeById(id int) *Node {
	start := L.Head
	for start != nil {
		if start.Album.ID == id {
			return start
		}
		start = start.next
	}
	return nil
}

func (L *List) ConvertToArray() []Album {
	albumSlice := make([]Album, 0)
	start := L.Head
	for start != nil {
		albumSlice = append(albumSlice, start.Album)
		start = start.next
	}
	return albumSlice
}

func (L *List) ConvertToReverseArray() []Album {
	albumSlice := make([]Album, 0)
	start := L.Tail
	for start != nil {
		albumSlice = append(albumSlice, start.Album)
		start = start.prev
	}
	return albumSlice
}
