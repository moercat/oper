package oper

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	squared := Map(numbers, func(x int) int { return x * x })
	expected := []int{1, 4, 9, 16, 25}

	if !reflect.DeepEqual(squared, expected) {
		t.Errorf("Map failed: expected %v, got %v", expected, squared)
	}
}

func TestFilter(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	expected := []int{2, 4}

	if !reflect.DeepEqual(evens, expected) {
		t.Errorf("Filter failed: expected %v, got %v", expected, evens)
	}
}

func TestAny(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	hasEven := Any(numbers, func(x int) bool { return x%2 == 0 })

	if !hasEven {
		t.Errorf("Any failed: expected true, got %v", hasEven)
	}

	allOdd := Any([]int{1, 3, 5}, func(x int) bool { return x%2 == 0 })
	if allOdd {
		t.Errorf("Any failed: expected false, got %v", allOdd)
	}
}

func TestAll(t *testing.T) {
	allPositive := All([]int{1, 2, 3, 4, 5}, func(x int) bool { return x > 0 })

	if !allPositive {
		t.Errorf("All failed: expected true, got %v", allPositive)
	}

	mixed := All([]int{1, 2, -3, 4, 5}, func(x int) bool { return x > 0 })
	if mixed {
		t.Errorf("All failed: expected false, got %v", mixed)
	}
}

func TestFind(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	value, index := Find(numbers, func(x int) bool { return x > 3 })

	if value != 4 || index != 3 {
		t.Errorf("Find failed: expected value 4 at index 3, got value %d at index %d", value, index)
	}

	_, index2 := Find(numbers, func(x int) bool { return x > 10 })
	if index2 != -1 {
		t.Errorf("Find failed: expected index -1, got %d", index2)
	}
}

func TestChunk(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	chunks := Chunk(numbers, 2)
	expected := [][]int{{1, 2}, {3, 4}, {5, 6}}

	if !reflect.DeepEqual(chunks, expected) {
		t.Errorf("Chunk failed: expected %v, got %v", expected, chunks)
	}

	// Test with remainder
	numbers2 := []int{1, 2, 3, 4, 5}
	chunks2 := Chunk(numbers2, 2)
	expected2 := [][]int{{1, 2}, {3, 4}, {5}}

	if !reflect.DeepEqual(chunks2, expected2) {
		t.Errorf("Chunk failed: expected %v, got %v", expected2, chunks2)
	}
}

func TestFlatten(t *testing.T) {
	nested := [][]int{{1, 2}, {3, 4}, {5}}
	flattened := Flatten(nested)
	expected := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(flattened, expected) {
		t.Errorf("Flatten failed: expected %v, got %v", expected, flattened)
	}
}

func TestContains(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	if !Contains(numbers, 3) {
		t.Errorf("Contains failed: expected true for element 3")
	}

	if Contains(numbers, 10) {
		t.Errorf("Contains failed: expected false for element 10")
	}
}

func TestReverse(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	reversed := Reverse(numbers)
	expected := []int{5, 4, 3, 2, 1}

	if !reflect.DeepEqual(reversed, expected) {
		t.Errorf("Reverse failed: expected %v, got %v", expected, reversed)
	}
}

func TestTake(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	taken := Take(numbers, 3)
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(taken, expected) {
		t.Errorf("Take failed: expected %v, got %v", expected, taken)
	}

	taken2 := Take(numbers, 10) // More than available
	expected2 := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(taken2, expected2) {
		t.Errorf("Take failed: expected %v, got %v", expected2, taken2)
	}
}

func TestDrop(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	dropped := Drop(numbers, 2)
	expected := []int{3, 4, 5}

	if !reflect.DeepEqual(dropped, expected) {
		t.Errorf("Drop failed: expected %v, got %v", expected, dropped)
	}

	dropped2 := Drop(numbers, 10) // More than available
	expected2 := []int{}

	if !reflect.DeepEqual(dropped2, expected2) {
		t.Errorf("Drop failed: expected %v, got %v", expected2, dropped2)
	}
}

func TestUniqBy(t *testing.T) {
	words := []string{"apple", "ant", "banana", "berry", "cat"}
	unique := UniqBy(words, func(s string) string { return string(s[0]) })
	expected := []string{"apple", "banana", "cat"} // 选取每个首字母的第一个单词

	if len(unique) != len(expected) {
		t.Errorf("UniqBy failed: expected length %d, got length %d", len(expected), len(unique))
	} else {
		// Check that all elements are as expected by their first letter
		for i, word := range expected {
			if string(word[0]) != string(unique[i][0]) {
				t.Errorf("UniqBy failed: expected %s at position %d, got %s", string(word[0]), i, string(unique[i][0]))
			}
		}
	}
}

func TestGroupBy(t *testing.T) {
	words := []string{"apple", "ant", "banana", "berry", "cat"}
	grouped := GroupBy(words, func(s string) rune { return rune(s[0]) })

	// Check that we have groups for 'a', 'b', 'c'
	if len(grouped) != 3 {
		t.Errorf("GroupBy failed: expected 3 groups, got %d", len(grouped))
	}

	if len(grouped['a']) != 2 || grouped['a'][0] != "apple" || grouped['a'][1] != "ant" {
		t.Errorf("GroupBy failed for 'a': got %v", grouped['a'])
	}

	if len(grouped['b']) != 2 || grouped['b'][0] != "banana" || grouped['b'][1] != "berry" {
		t.Errorf("GroupBy failed for 'b': got %v", grouped['b'])
	}

	if len(grouped['c']) != 1 || grouped['c'][0] != "cat" {
		t.Errorf("GroupBy failed for 'c': got %v", grouped['c'])
	}
}
