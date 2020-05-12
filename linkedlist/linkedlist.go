package linkedlist
import (
	 "fmt"
	"personell/datastructure-algorithem/dsstack"
	"runtime/debug"
)
//
// ReverseLinkedList
//
func (this *LL) ReverseLinkedList() {
	
	var prev *Node
	curr := this.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	this.Head = prev
}

func (this *LL) FindMid()int {

	pt1 := this.Head
	pt2 := this.Head.Next
	for pt2.Next != nil && pt2.Next.Next != nil {
		pt1 = pt1.Next
		pt2 = pt2.Next.Next
	}
	return pt1.Next.Value
}

func (this *LL) FindNth(n int)int {
	pt1 := this.Head
	pt2 := this.Head
	var count int
	for pt2 != nil {
		if count > n {
			pt2 = pt2.Next
			pt1 = pt1.Next
		}else {
			count = count+1
			pt2 = pt2.Next
		}
	}
	if count < n {
		//error return
	}
	return pt1.Value
}

func (this *LL)RemoveLoop() {
	pt1 := this.Head
	pt2 := this.Head.Next

	for pt1 != pt2 && pt2.Next != nil {
		//detech loop
		pt1 = pt1.Next
		pt2 = pt2.Next.Next
	}
	//we found the node which is circular
	//now we fix pt2 and put pt1 to head
	//move pt1 and pt2 till it's not equal to pt1 
	//other wise come to same place pt2
	//the time both mean it;s the cirucular point then just break it

	pt1  = this.Head
	loopNode := pt2
	for {
		if pt1 == pt2.Next {
			break
		}else {
			pt1 = pt1.Next
		}
		for pt2.Next != loopNode && pt1 != pt2.Next && pt2.Next.Next != pt2 {
			pt2 = pt2.Next
		}
	}
	pt2.Next = nil	
}


func (this *LL)ReverseTillK(k int) {

	var prev *Node
	curr := this.Head
	var count int
	var next *Node
	for count < k && curr != nil  {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count = count+1
	}
	this.Head.Next = next
	this.Head = prev
}


func Add2Numbers(l1 *LL,l2 *LL)*LL {
	defer func() {
		if r := recover(); r!= nil {
			debug.PrintStack()
		}
	}()
	firstList := l1.Head
	secondList := l2.Head
	var carry int
	var result LL
	for firstList != nil || secondList != nil {
		var val1,val2 int
		if firstList != nil {
			val1 = firstList.Value
			firstList = firstList.Next
		}
		if secondList != nil {
			val2 = secondList.Value
			secondList = secondList.Next
		}
		valu3 := carry + val1+val2
		if valu3 >=10 {
			carry = 1
		}else {
			carry = 0
		}
		valu3 = valu3%10
		var node Node
		if result.Head == nil {
			node = Node{
				Value : valu3,
				Next : nil,
			}
		}else {
			node = Node{
				Value : valu3,
				Next : result.Head,
			}
		}
		result.Head = &node
	}
	return &result
}

//
// CheckPalindrome
//
func(this *LL) CheckPalindrome()bool {

	tmp := this.Head
	stack := dsstack.InitialiseStack()
	for tmp != nil {
		stack.Push(tmp.Value)
		tmp = tmp.Next
	}
	for this.Head != nil {
		if stack.Top() == this.Head.Value {
			stack.Pop()
			this.Head = this.Head.Next
		} else {
			return false
		}
	}
	return true
}


func (this *LL)Insert(data int) {

	if this.Head == nil {
		newNode := Node{
			Value : data,
			Next : nil,
		}
		this.Head = &newNode
	}else {
		newNode := Node{
			Value : data,
			Next : this.Head,
		}
		this.Head = &newNode
	}
}

func (this *LL)TraverseLinkedList()[]int {

	a := make([]int,0)
	t := this.Head
	for t != nil {
		a = append(a,t.Value)
		t = t.Next
	}
	return a
}

func (this *LL)ReArrange() {

	pt1 := this.Head
	pt2 := this.Head.Next
	pt3 := this.Head
	//find mid first
	for pt2.Next != nil && pt2.Next.Next != nil {
		pt1 = pt1.Next
		pt2 = pt2.Next.Next
	}
	var s2  LL
	s2.Head = pt1.Next
	fmt.Printf("dddd %d",pt1.Next.Value)
	s2.ReverseLinkedList()
	s3 := s2.Head
	pt1.Next = nil 

	tmp := pt3.Next
	for tmp != nil {
		fmt.Printf("seeee value %d",s3.Value)
		pt3.Next = s3
		pt3.Next.Next = tmp
		tmp = tmp.Next
		s3 = s3.Next
		pt3 = pt3.Next.Next
	}
	fmt.Printf("s3 data %d",s3.Value)
}

