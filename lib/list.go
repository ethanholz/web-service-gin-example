package lib

type Node struct {
    Album Album
    next *Node
}

type List struct {
    Head *Node
}

func (L *List) Insert(album Album){
    list := &Node {
        Album: album,
        next: L.Head,
    }
    L.Head = list
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

func (L *List) ConvertToList() []Album {
    var albumArray []Album
    start := L.Head
    for start != nil {
        albumArray = append(albumArray, start.Album)
        start = start.next
    }
    return albumArray
}
