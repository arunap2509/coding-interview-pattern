package datastructure

type Deque struct { 
    items []int
} 

func NewDeque() *Deque { 
    return new(Deque) 
}

func (s *Deque) PushFront(item int) { 
    temp := []int{item} 
    s.items = append(temp, s.items...) 
} 

func (s *Deque) PushBack(item int) { 
    s.items = append(s.items, item)
} 

func (s *Deque) PopFront() int { 
    defer func() { 
        s.items = s.items[1:] 
    }()
    return s.items[0] 
} 

func (s *Deque) PopBack() int { 
    i := len(s.items) - 1 
    defer func() { 
        s.items = s.items[:i]
    }()
    return s.items[i] 
} 

func (s *Deque) Front() int { 
    return s.items[0] 
} 

func (s *Deque) Back() int { 
    return s.items[len(s.items) - 1] 
} 

func (s *Deque) Empty() bool { 
    if len(s.items) == 0 { 
        return true 
    } 
    return false 
} 

func (s *Deque) Len() int { 
    return len(s.items)
} 