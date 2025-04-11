package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "xd")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stact", func(t *testing.T) {
		myStackOfInts := NewStack[int]()

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(465)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 465)

		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		first, _ := myStackOfInts.Pop()
		second, _ := myStackOfInts.Pop()
		AssertEqual(t, first+second, 3)

	})
	t.Run("string stact", func(t *testing.T) {
		myStackOfInts := NewStack[string]()

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push("123")
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push("465")
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, "465")

		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, "123")

		AssertTrue(t, myStackOfInts.IsEmpty())
	})

}
