package dsstack


//using stack we will do that
// idea is we will push the element in stack and compare it with upcoming value if the value of stack is low
// we are create a pair and pop and check for next stack value and if stack empty just push 
//current value in stack and increase i to check further pair
//if at last stack has value it's pair should be -1
func FindNextGreater(arr []int) [][]int{
	output := make([][]int,0)
	st := InitialiseStack()
	if arr == nil || len(arr)==0 {
		return nil
	}
	st.Push(arr[0])
	for i:=1;i<len(arr);i++ {
		for st.IsEmpty() == false && st.Top() < arr[i] {
			t := []int{st.Top(),arr[i]}
			output = append(output,t)
			st.Pop()
		}
		st.Push(arr[i])
	}
	for st.IsEmpty() == false {
		t := []int{st.Top(),-1}
		output = append(output,t)
		st.Pop()
	}
	return output
}


func InitialiseStack() *Stack{
	st := Stack{}
	return &st
}

func(this *Stack)Push(a int) {
	this.Data = append(this.Data,a)
}

func (this *Stack)Pop() {
	this.Data = this.Data[:len(this.Data)-1]
}

func (this *Stack)Top() int{
	return this.Data[len(this.Data)-1]
}

func (this *Stack)IsEmpty()bool {
	if len(this.Data)==0 {
		return true
	}
	return false
}