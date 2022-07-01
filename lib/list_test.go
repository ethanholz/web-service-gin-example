package lib

import (
	"fmt"
	"testing"
)

func TestInsertStart(t *testing.T) {
	var testList List
	album := Album{
		Title:  "Test Album",
		Artist: "Artist 1",
		ID:     1,
		Price:  1.50,
	}
	album2 := Album{
		Title:  "Test Album 2",
		Artist: "Artist 2",
		ID:     2,
		Price:  1.50,
	}
	testList.InsertStart(album)
	if testList.Head.Album != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}
	if testList.Tail.Album != album {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album)
	}
	testList.InsertStart(album2)
	if testList.Head.Album != album2 {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album2)
	}
	if testList.Tail.Album != album {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album)
	}
}

func TestInsert(t *testing.T) {
	var testList List
	album := Album{
		Title:  "Test Album",
		Artist: "Artist 1",
		ID:     1,
		Price:  1.50,
	}
	album2 := Album{
		Title:  "Test Album 2",
		Artist: "Artist 2",
		ID:     2,
		Price:  1.50,
	}
	testList.Insert(album)
	if testList.Head.Album != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}
	if testList.Tail.Album != album {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album)
	}
	testList.Insert(album2)
	if testList.Tail.Album != album2 {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album2)
	}
	if testList.Head.Album != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}
}

func TestRemoveTail(t *testing.T) {
	var testList List
	album := Album{
		Title:  "Test Album",
		Artist: "Artist 1",
		ID:     1,
		Price:  1.50,
	}
	album2 := Album{
		Title:  "Test Album 2",
		Artist: "Artist 2",
		ID:     2,
		Price:  1.50,
	}
	_, err := testList.RemoveTail()
	if err == nil {
		t.Error("unable to handle empty list")
	}
	testList.Insert(album)
	val, _ := testList.RemoveTail()
	if testList.Head != nil && testList.Tail != nil {
		t.Error("unable to remove last element")
	}
	if *val != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}

	testList.Insert(album)
	if testList.Head.Album != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}
	if testList.Tail.Album != album {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album)
	}
	testList.Insert(album2)
	if testList.Tail.Album != album2 {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album2)
	}
	if testList.Head.Album != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}

	retAlbum, err := testList.RemoveTail()
	if err != nil && retAlbum == nil {
		t.Errorf("Unable to remove correctly: %s", err)
	}
	if retAlbum.String() != album2.String() {
		t.Errorf("got %s, wanted %s", retAlbum.String(), album2.String())
	}
	if testList.Head.Album != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}
	if testList.Tail.Album != album {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album)
	}
}

func TestRemoveHead(t *testing.T) {
	var testList List
	album := Album{
		Title:  "Test Album",
		Artist: "Artist 1",
		ID:     1,
		Price:  1.50,
	}
	album2 := Album{
		Title:  "Test Album 2",
		Artist: "Artist 2",
		ID:     2,
		Price:  1.50,
	}
	_, err := testList.RemoveHead()
	if err == nil {
		t.Error("unable to handle empty list")
	}

	testList.Insert(album)
	val, _ := testList.RemoveHead()
	if testList.Head != nil && testList.Tail != nil {
		t.Error("unable to remove last element")
	}
	if *val != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}

	testList.Insert(album)
	if testList.Head.Album != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}
	if testList.Tail.Album != album {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album)
	}
	testList.Insert(album2)
	if testList.Tail.Album != album2 {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album2)
	}
	if testList.Head.Album != album {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album)
	}

	retAlbum, err := testList.RemoveHead()
	if err != nil && retAlbum == nil {
		t.Errorf("Unable to remove correctly: %s", err)
	}
	if retAlbum.String() != album.String() {
		t.Errorf("got %s, wanted %s", retAlbum.String(), album.String())
	}
	if testList.Head.Album != album2 {
		t.Errorf("got %s, wanted %s", testList.Head.Album, album2)
	}
	if testList.Tail.Album != album2 {
		t.Errorf("got %s, wanted %s", testList.Tail.Album, album2)
	}
}

func TestGetNodeById(t *testing.T) {
	var testList List
	album := Album{
		Title:  "Test Album",
		Artist: "Artist 1",
		ID:     1,
		Price:  1.50,
	}
	album2 := Album{
		Title:  "Test Album 2",
		Artist: "Artist 2",
		ID:     2,
		Price:  1.50,
	}
	testList.Insert(album)
	testList.Insert(album2)
	node := testList.GetNodeById(1)
	if node == nil {
		t.Error("got nil, expected 1")
	}
	node = testList.GetNodeById(3)
	if node != nil {
		t.Errorf("got %d, expected nil", node.Album.ID)
	}
}

func TestConvertToArray(t *testing.T) {
	var testList List
	album := Album{
		Title:  "Test Album",
		Artist: "Artist 1",
		ID:     1,
		Price:  1.50,
	}
	album2 := Album{
		Title:  "Test Album 2",
		Artist: "Artist 2",
		ID:     2,
		Price:  1.50,
	}
	albums := []Album{album, album2}

	compAlbums := testList.ConvertToArray()
	if len(compAlbums) != 0 {
		t.Errorf("got %d, expected 0", len(compAlbums))
	}

	testList.Insert(album)
	testList.Insert(album2)
	compAlbums = testList.ConvertToArray()
	if len(compAlbums) == 0 {
		t.Errorf("got %d, expected 2", len(compAlbums))
		return
	}
	for i, album := range compAlbums {
		fmt.Printf("%d", i)
		if album != albums[i] {
			t.Errorf("got %d, expected %d", album.ID, albums[i].ID)
		}
	}
}
func TestConvertToReverseArray(t *testing.T) {
	var testList List
	album := Album{
		Title:  "Test Album",
		Artist: "Artist 1",
		ID:     1,
		Price:  1.50,
	}
	album2 := Album{
		Title:  "Test Album 2",
		Artist: "Artist 2",
		ID:     2,
		Price:  1.50,
	}
	albums := []Album{album2, album}

	compAlbums := testList.ConvertToReverseArray()
	if len(compAlbums) != 0 {
		t.Errorf("got %d, expected 0", len(compAlbums))
	}

	testList.Insert(album)
	testList.Insert(album2)
	compAlbums = testList.ConvertToReverseArray()
	if len(compAlbums) == 0 {
		t.Errorf("got %d, expected 2", len(compAlbums))
		return
	}
	for i, album := range compAlbums {
		fmt.Printf("%d", i)
		if album != albums[i] {
			t.Errorf("got %d, expected %d", album.ID, albums[i].ID)
		}
	}
}
