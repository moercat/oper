package oper

import (
	"reflect"
	"testing"
)

func TestStringUtilities(t *testing.T) {
	// Test Capitalize
	if result := Capitalize("hello"); result != "Hello" {
		t.Errorf("Capitalize failed: expected 'Hello', got '%s'", result)
	}

	if result := Capitalize("hELLO"); result != "Hello" {
		t.Errorf("Capitalize failed: expected 'Hello', got '%s'", result)
	}

	// Test CamelCase
	if result := CamelCase("hello world"); result != "helloWorld" {
		t.Errorf("CamelCase failed: expected 'helloWorld', got '%s'", result)
	}

	if result := CamelCase("hello_world-test"); result != "helloWorldTest" {
		t.Errorf("CamelCase failed: expected 'helloWorldTest', got '%s'", result)
	}

	// Test KebabCase
	if result := KebabCase("HelloWorld"); result != "hello-world" {
		t.Errorf("KebabCase failed: expected 'hello-world', got '%s'", result)
	}

	if result := KebabCase("hello_world test"); result != "hello-world-test" {
		t.Errorf("KebabCase failed: expected 'hello-world-test', got '%s'", result)
	}

	// Test SnakeCase
	if result := SnakeCase("HelloWorld"); result != "hello_world" {
		t.Errorf("SnakeCase failed: expected 'hello_world', got '%s'", result)
	}

	if result := SnakeCase("hello-world test"); result != "hello_world_test" {
		t.Errorf("SnakeCase failed: expected 'hello_world_test', got '%s'", result)
	}
}

func TestSetOperations(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	// Test Difference
	diff := Difference(slice1, slice2)
	expectedDiff := []int{1, 2}
	if !reflect.DeepEqual(diff, expectedDiff) {
		t.Errorf("Difference failed: expected %v, got %v", expectedDiff, diff)
	}

	// Test Intersection
	inter := Intersection(slice1, slice2)
	expectedInter := []int{3, 4, 5}
	if !reflect.DeepEqual(inter, expectedInter) {
		t.Errorf("Intersection failed: expected %v, got %v", expectedInter, inter)
	}

	// Test UnionSlice
	union := UnionSlice(slice1, slice2)
	expectedUnion := []int{1, 2, 3, 4, 5, 6, 7}
	if len(union) != len(expectedUnion) {
		t.Errorf("UnionSlice failed: expected length %d, got %d", len(expectedUnion), len(union))
	}

	// Test SymmetricDifference
	symDiff := SymmetricDifference(slice1, slice2)
	expectedSymDiff := []int{1, 2, 6, 7}
	if len(symDiff) != len(expectedSymDiff) {
		t.Errorf("SymmetricDifference failed: expected length %d, got %d", len(expectedSymDiff), len(symDiff))
	}
}

func TestNumericUtilities(t *testing.T) {
	// Test Range
	rangeResult := Range(1, 5, 1)
	expectedRange := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(rangeResult, expectedRange) {
		t.Errorf("Range failed: expected %v, got %v", expectedRange, rangeResult)
	}

	// Test Sum
	if sum := Sum([]int{1, 2, 3, 4}); sum != 10 {
		t.Errorf("Sum failed: expected 10, got %d", sum)
	}

	// Test Avg
	if avg := Avg([]int{1, 2, 3, 4}); avg != 2.5 {
		t.Errorf("Avg failed: expected 2.5, got %f", avg)
	}

	// Test Min
	if min := Min([]int{3, 1, 4, 1, 5}); min != 1 {
		t.Errorf("Min failed: expected 1, got %d", min)
	}

	// Test Max
	if max := Max([]int{3, 1, 4, 1, 5}); max != 5 {
		t.Errorf("Max failed: expected 5, got %d", max)
	}

	// Test Clamp
	if clamped := Clamp(10, 1, 5); clamped != 5 {
		t.Errorf("Clamp failed: expected 5, got %d", clamped)
	}
	if clamped := Clamp(0, 1, 5); clamped != 1 {
		t.Errorf("Clamp failed: expected 1, got %d", clamped)
	}
	if clamped := Clamp(3, 1, 5); clamped != 3 {
		t.Errorf("Clamp failed: expected 3, got %d", clamped)
	}
}

func TestArrayUtilities(t *testing.T) {
	slice := []int{0, 1, 2, 0, 3}

	// Test Compact
	compactResult := Compact(slice)
	expectedCompact := []int{1, 2, 3}
	if !reflect.DeepEqual(compactResult, expectedCompact) {
		t.Errorf("Compact failed: expected %v, got %v", expectedCompact, compactResult)
	}

	// Test Uniq
	dupSlice := []int{1, 2, 2, 3, 3, 4}
	uniqResult := Uniq(dupSlice)
	expectedUniq := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(uniqResult, expectedUniq) {
		t.Errorf("Uniq failed: expected %v, got %v", expectedUniq, uniqResult)
	}

	// Test Initial
	initialResult := Initial([]int{1, 2, 3, 4})
	expectedInitial := []int{1, 2, 3}
	if !reflect.DeepEqual(initialResult, expectedInitial) {
		t.Errorf("Initial failed: expected %v, got %v", expectedInitial, initialResult)
	}

	// Test Tail
	tailResult := Tail([]int{1, 2, 3, 4})
	expectedTail := []int{2, 3, 4}
	if !reflect.DeepEqual(tailResult, expectedTail) {
		t.Errorf("Tail failed: expected %v, got %v", expectedTail, tailResult)
	}

	// Test TakeRight
	takeRightResult := TakeRight([]int{1, 2, 3, 4, 5}, 2)
	expectedTakeRight := []int{4, 5}
	if !reflect.DeepEqual(takeRightResult, expectedTakeRight) {
		t.Errorf("TakeRight failed: expected %v, got %v", expectedTakeRight, takeRightResult)
	}

	// Test DropRight
	dropRightResult := DropRight([]int{1, 2, 3, 4, 5}, 2)
	expectedDropRight := []int{1, 2, 3}
	if !reflect.DeepEqual(dropRightResult, expectedDropRight) {
		t.Errorf("DropRight failed: expected %v, got %v", expectedDropRight, dropRightResult)
	}
}

func TestFindObject(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}

	// Test FindLast - "cherry" is the last one with len > 5 (both "banana" and "cherry" have len > 5)
	result, found := FindLast(slice, func(item string) bool {
		return len(item) > 5
	})

	if !found || result != "cherry" { // "cherry" and "banana" both have len > 5, but "cherry" is the last one
		t.Errorf("FindLast failed: expected 'cherry', found: %t, result: %s", found, result)
	}

	// Correct example: find the last item that starts with 'b'
	result, found = FindLast(slice, func(item string) bool {
		return item[0] == 'b'
	})

	if !found || result != "banana" {
		t.Errorf("FindLast failed: expected 'banana', found: %t, result: %s", found, result)
	}
}

func TestObjectHelpers(t *testing.T) {
	obj := map[string]interface{}{
		"name": "John",
		"age":  30,
		"address": map[string]interface{}{
			"city": "New York",
			"zip":  "10001",
		},
		"hobbies": []string{"reading", "swimming"},
	}

	// Test Get
	if result := Get(obj, "name", "default").(string); result != "John" {
		t.Errorf("Get failed: expected 'John', got '%s'", result)
	}

	if result := Get(obj, "address.city", "default").(string); result != "New York" {
		t.Errorf("Get failed: expected 'New York', got '%s'", result)
	}

	if result := Get(obj, "address.country", "default").(string); result != "default" {
		t.Errorf("Get failed: expected 'default', got '%s'", result)
	}

	// Test Pick
	picked := Pick(obj, []string{"name", "age"})
	expectedPicked := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	if !reflect.DeepEqual(picked, expectedPicked) {
		t.Errorf("Pick failed: expected %v, got %v", expectedPicked, picked)
	}

	// Test Omit
	omitted := Omit(obj, []string{"age", "hobbies"})

	// For omitted test, we'll just check that 'age' and 'hobbies' are not present
	if _, exists := omitted["age"]; exists {
		t.Errorf("Omit failed: 'age' should not be present in result")
	}
	if _, exists := omitted["hobbies"]; exists {
		t.Errorf("Omit failed: 'hobbies' should not be present in result")
	}
	if _, exists := omitted["name"]; !exists {
		t.Errorf("Omit failed: 'name' should be present in result")
	}
}
