package main

type IItem interface {
	GetID() string
}

type Item struct {
	ID string
}

func (i *Item) GetID() string {
	return i.ID
}

func main() {
	f := []Item{}
	for i := 0; i < 10; i++ {
		f = append(f, Item{ID: "nice"})
	}
	notnice(f)
}

func notnice(in interface{}) {
	_, ok := in.([]any)
	println(ok)
}
