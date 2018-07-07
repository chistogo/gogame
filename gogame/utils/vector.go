package utils


type VectorF []float32
type VectorI []int32

func NewVectorI(dims ...int32) VectorI{

	coordinate := []int32{}
	for _, dim := range dims {
		coordinate = append(coordinate,dim)
	}
	return coordinate
}

func NewVectorF(dims ...float32) VectorF{

	coordinate := []float32{}
	for _, dim := range dims {
		coordinate = append(coordinate,dim)
	}
	return coordinate

}