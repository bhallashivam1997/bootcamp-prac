package main

type Stack []et
type et struct{
	value byte
	left,right *et
}


func isOperator(c byte) bool{
	if c=='+' || c=='-' || c=='*' || c=='/'{
		return true
	}
	return false
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str et) {
	*s = append(*s, str)
}

func (s *Stack) Pop(){
	if s.IsEmpty() {
	} else {
		index := len(*s) - 1
		*s = (*s)[:index]
	}
}

func (s *Stack) Top() et{
	element := (*s)[len(*s) - 1]
	return element
}

func buildExpressionTree(str string) *et{
	var stack Stack
	var t,t1,t2 *et

	for i:=0 ; i<len(str) ; i++{
		if !isOperator(str[i]){
			var tmpt= et{str[i],nil,nil}
			t=&tmpt
			stack.Push(*t)
		}else{
			var tmpt= et{str[i],nil,nil}
			t= &tmpt
			*t1 = stack.Top();
			stack.Pop();
			*t2 = stack.Top();
			stack.Pop();


			t.right = t1
			t.left = t2
			stack.Push(*t);
		}
	}
	*t = stack.Top();
	stack.Pop();
	return t
}



func main() {
	var expression string ="a+b*c/d"
	buildExpressionTree(expression)
}
