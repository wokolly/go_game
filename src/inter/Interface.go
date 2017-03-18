package inter

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type SortableStrings [3]string

func (self SortableStrings) Len() int {
	return len(self)
}

func (self SortableStrings) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self SortableStrings) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}
