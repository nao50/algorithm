package main

import "fmt"

type Element struct {
	item int
	next *Element
}

type List struct {
	root *Element
}

func newElement(x int, ep *Element) *Element {
	return &Element{
		item: x,
		next: ep,
	}
}

func newList() *List {
	return &List{
		root: new(Element),
	}
}

// データの挿入
func (list *List) insertElement(n, x int) bool {
	ep := list.root.getElement(n - 1)
	if ep == nil {
		return false
	}
	ep.next = newElement(x, ep.next)
	return true
}

// n 番目のセルを求める
func (ep *Element) getElement(n int) *Element {
	i := -1
	for ep != nil {
		if i == n {
			return ep
		}
		i++
		ep = ep.next
	}
	return nil
}

// n 番目の要素を求める
func (lst *List) nth(n int) (int, bool) {
	ep := lst.root.getElement(n)
	if ep == nil {
		return 0, false
	}
	return ep.item, true
}

// データの削除
func (lst *List) deleteNth(n int) bool {
	ep := lst.root.getElement(n - 1)
	if ep == nil || ep.next == nil {
		return false
	}
	ep.next = ep.next.next
	return true
}

// 連結リストは空か？
func (lst *List) isEmpty() bool {
	return lst.root.next == nil
}

// 表示
func (lst *List) printList() {
	ep := lst.root.next
	for ; ep != nil; ep = ep.next {
		fmt.Print(ep.item, " ")
	}
	fmt.Println("")
}

func main() {
	a := newList()
	for i := 0; i < 4; i++ {
		// fmt.Println(a.insertElement(i, i))
		a.insertElement(i, i)
	}

	a.insertElement(4, 4)
	a.insertElement(6, 6)

	// a.printList()
	for i := 0; i < 5; i++ {
		n, ok := a.nth(i)
		fmt.Println(n, ok)
	}
	// for !a.isEmpty() {
	// 	a.deleteNth(0)
	// 	a.printList()
	// }
}
