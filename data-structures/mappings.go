package datastructures

type Mapping struct {
	Src     int
	Src_max int
	Dst     int
	Dst_max int
	Count   int
}
type MappingList []Mapping

func (m Mapping) Contains(v int) bool { return (v >= m.Src && v < m.Src_max) }
func (m Mapping) Get(v int) int       { return (v - m.Src) + m.Dst }

//
// type MappingNode struct {
// 	mapping Mapping
// 	prev    *Mapping
// 	next    *Mapping
// }
// type MappingLinkedList []MappingNode
