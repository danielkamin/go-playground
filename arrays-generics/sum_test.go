package arraysgenerics

import (
	"strings"
	"testing"
)

type Person struct {
	Name string
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
	t.Run("concat string", func(t *testing.T) {
		concat := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"1", "2", "3"}, concat, ""), "123")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		firsEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firsEvenNumber, 2)
	})

	t.Run("find the best dev", func(t *testing.T) {
		people := []Person{
			{Name: "ME"},
			{Name: "YOU"},
			{Name: "HIM"},
		}
		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "E")
		})
		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "ME"})
	})
}
