// Copyright (C) 2014 Makoto Hiroi
// http://www.nct9.ne.jp/m_hiroi/golang/abcgo07.html
package main

import "fmt"

// セル
type Cell struct {
	item int
	next *Cell
}

// 連結リスト
type List struct {
	top *Cell
}

// セルの生成
func newCell(x int, cp *Cell) *Cell {
	newcp := new(Cell)
	newcp.item, newcp.next = x, cp
	return newcp
}

// 連結リストの生成
func newList() *List {
	lst := new(List)
	lst.top = new(Cell)
	return lst
}

// n 番目のセルを求める
func (cp *Cell) nthCell(n int) *Cell {
	i := -1
	for cp != nil {
		if i == n {
			return cp
		}
		i++
		cp = cp.next
	}
	return nil
}

// n 番目の要素を求める
func (lst *List) nth(n int) (int, bool) {
	cp := lst.top.nthCell(n)
	if cp == nil {
		return 0, false
	}
	return cp.item, true
}

// データの挿入
func (lst *List) insertNth(n, x int) bool {
	cp := lst.top.nthCell(n - 1)
	if cp == nil {
		return false
	}
	cp.next = newCell(x, cp.next)
	return true
}

// データの削除
func (lst *List) deleteNth(n int) bool {
	cp := lst.top.nthCell(n - 1)
	if cp == nil || cp.next == nil {
		return false
	}
	cp.next = cp.next.next
	return true
}

// 連結リストは空か？
func (lst *List) isEmpty() bool {
	return lst.top.next == nil
}

// 表示
func (lst *List) printList() {
	cp := lst.top.next
	for ; cp != nil; cp = cp.next {
		fmt.Print(cp.item, " ")
	}
	fmt.Println("")
}

func main() {
	a := newList()
	for i := 0; i < 4; i++ {
		// fmt.Println(a.insertNth(i, i))
		a.insertNth(i, i)
	}

	a.insertNth(4, 4)
	a.insertNth(6, 6)

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
