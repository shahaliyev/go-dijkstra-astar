package main

// Edge type
type Edge struct {
	to int
	distance float64
}

// Less function for the Edge type priority queue
func (i Edge) Less(j QItem) bool {
	return i.distance < j.(Edge).distance
}
