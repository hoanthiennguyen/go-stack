package stack

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)

	if val, err := s.Pop(); err != nil || val != 2 {
		t.Errorf("Expected 2, got %v (err: %v)", val, err)
	}

	if val, err := s.Pop(); err != nil || val != 1 {
		t.Errorf("Expected 1, got %v (err: %v)", val, err)
	}

	if _, err := s.Pop(); err != errStackEmpty {
		t.Errorf("Expected errStackEmpty, got %v", err)
	}
}

func TestStackPeek(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)

	if val, err := s.Peek(); err != nil || val != 2 {
		t.Errorf("Expected 2, got %v (err: %v)", val, err)
	}

	s.Pop()

	if val, err := s.Peek(); err != nil || val != 1 {
		t.Errorf("Expected 1, got %v (err: %v)", val, err)
	}

	s.Pop()

	if _, err := s.Peek(); err != errStackEmpty {
		t.Errorf("Expected errStackEmpty, got %v", err)
	}
}

func TestStackLen(t *testing.T) {
	s := New[int]()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}

	s.Push(1)
	s.Push(2)
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	s.Pop()
	if s.Len() != 1 {
		t.Errorf("Expected length 1, got %d", s.Len())
	}

	s.Pop()
	if s.Len() != 0 {
		t.Errorf("Expected length 0, got %d", s.Len())
	}
}

func TestStackIsEmpty(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("Expected stack to not be empty")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}
