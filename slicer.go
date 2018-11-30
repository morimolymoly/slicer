package slicer

// Index ... Index
type Index struct {
	Start int
	End   int
}

// GetSliceIndex ... Get Index Each BufferSize
func GetSliceIndex(sliceLength, BufferSize int) []Index {
	var ret = []Index{}
	for i := 0; i < sliceLength; i += BufferSize {
		end := i + BufferSize
		if sliceLength < end {
			end = sliceLength
		}
		ret = append(ret, Index{Start: i, End: end})
	}
	return ret
}
