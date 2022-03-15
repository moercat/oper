package list

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestOperand(t *testing.T) {

	ins := []int{
		1, 2, 3,
	}

	l := List{
		ins,
	}
	l.Pop(1)
	fmt.Println(l[0])
	l.Insert(1,3)
	fmt.Println(l[0])
	inSlice := cast.ToIntSlice(l[0])
	fmt.Println(inSlice[0])

}
