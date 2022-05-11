package stack

import "testing"

func TestNew(t *testing.T) {
	v := []int{7, 2, 5}
	s := New(v...)

	if s.size != 3 {
		t.Errorf("Expected size 3, got %d", s.size)
	}

	n := s.head
	var i uint
	for n != nil {
		if n.value != v[i] {
			t.Errorf("Expected %d, got %d", v[i], n.value)
		}

		n = n.next
		i++
	}
}

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(2)
	s.Push(7)

	if s.size != 2 {
		t.Errorf("Expected size 2, got %d", s.size)
	}

	n := s.head
	if n.value != 7 {
		t.Errorf("Expected 7, got %d", n.value)
	}

	n = n.next
	if n.value != 2 {
		t.Errorf("Expected 2, got %d", n.value)
	}
}

func TestPop(t *testing.T) {
	s := New(7, 2, 5)

	if s.Pop() != 7 {
		t.Errorf("Expected 7, got %d", s.Pop())
	}

	if s.size != 2 {
		t.Errorf("Expected size 2, got %d", s.size)
	}

	if s.Pop() != 2 {
		t.Errorf("Expected 2, got %d", s.Pop())
	}

	if s.size != 1 {
		t.Errorf("Expected size 1, got %d", s.size)
	}

	if s.head.value != 5 {
		t.Errorf("Expected 5, got %d", s.head.value)
	}
}

func TestGetSize(t *testing.T) {
	s := New(7, 2, 5)

	if s.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", s.GetSize())
	}
}
