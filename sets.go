package sets

// Set represents a collection of unique elements.
type Set[T comparable] struct {
	elements map[T]struct{}
}

// New creates and returns a new empty set.
func New[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// Add inserts an element into the set.
func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

// Remove deletes an element from the set.
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// Contains checks if an element is in the set.
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// Union returns a new set that is the union of two sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := New[T]()
	for elem := range s.elements {
		result.Add(elem)
	}
	for elem := range other.elements {
		result.Add(elem)
	}
	return result
}

// Intersection returns a new set that is the intersection of two sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()
	for elem := range s.elements {
		if other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Difference returns a new set that is the difference of two sets.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()
	for elem := range s.elements {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}
	return result
}

// IsSubset checks if the set is a subset of another set.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for elem := range s.elements {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// ToArray converts the set to a slice.
func (s *Set[T]) ToArray() []T {
	arr := make([]T, 0, len(s.elements))
	for elem := range s.elements {
		arr = append(arr, elem)
	}
	return arr
}
