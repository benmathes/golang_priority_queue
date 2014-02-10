package main

import (
  "container/heap"
  "fmt"
)

// something we manage in a queue -- type definition
type Item struct {
  value string // the item value
  priority int // 
  index int // index of item in the underlying heap. why duplicate with array index? for the heap.
}


// type declaration. Implements heap.Interface and holds Items (array of Item pointers)
type PriorityQueue []*Item 


// the length of the priority queue pq, (pq.Len()) is the len of pq()
func (pq PriorityQueue) Len() int { return len(pq) } 


// given indexes i and j into a priority queue, tells us which is higher priority
func (pq PriorityQueue) Less(i,j int) bool {
  // pop should give us the highest priority
  return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i,j int) {
  pq[i], pq[j] = pq[j], pq[i] //update in array
  pq[i].index = i // update in heap
  pq[j].index = j // update in heap
}
  


// push x onto the pqueue
func (pq *PriorityQueue) Push(x interface{}) {
  n := len(*pq) // length of the pointer
  item := x.(*Item)
  item.index = n
  *pq = append(*pq, item)
}


// pop the highest-priority item
func (pq *PriorityQueue) Pop() interface{} {
  old := *pq
  length := len(old)
  item := old[length-1] // last item in list, which is highest priority
  item.index = -1 // Q safety
  *pq = old[0:length-1] // queue is now the slice upto but not including the popped element. GC will clean up the last Item.
  return item
}


func (pq *PriorityQueue) update(item *Item, value string, priority int) {
  heap.Remove(pq, item.index)
  item.value = value
  item.priority = priority
  heap.Push(pq, item)
}


// play with the PQ
func main() {
  items := map[string]int{
    "a": 1,
    "b": 10,
    "c": 100,
    "d": 101,
  }

  // create and insert
  pq := &PriorityQueue{}
  heap.Init(pq)
  for value, priority := range items {
    heap.Push(pq, &Item{ value: value, priority: priority })
  }


  new_item := &Item{
    value: "new_item", 
    priority: 5,
  }
  heap.Push(pq, new_item)

  pq.update(new_item, new_item.value, 11)

  
  // take them out in order
  for pq.Len() > 0 {
    item := heap.Pop(pq).(*Item)
    fmt.Printf("%.2d:%s ", item.priority, item.value)
  }

}