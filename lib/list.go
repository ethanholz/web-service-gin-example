package lib

type Node struct {
	Album Album
	next  *Node
	prev  *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func (L *List) RemoveTail() {
	lastNode := L.Tail
	if L.Tail.prev == nil && L.Tail.next == nil {
		L.Tail = nil
		L.Head = nil
		return
	}
	newTail := lastNode.prev
	newTail.next = nil
	L.Tail = newTail
}

func (L *List) InsertStart(album Album) {
	newNode := &Node{
		Album: album,
		next:  nil,
		prev:  nil,
	}

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

func (L *List) GetLength() int {
	length := 0
	start := L.Head
	length++
	for start != nil {
		start = start.next
		length++
	}
	return length
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
	var albumArray []Album
	start := L.Head
	for start != nil {
		albumArray = append(albumArray, start.Album)
		start = start.next
	}
	return albumArray
}

func (L *List) ConvertToReverseArray() []Album {
	var albumArray []Album
	start := L.Tail
	for start != nil {
		albumArray = append(albumArray, start.Album)
		start = start.prev
	}
	return albumArray
}
