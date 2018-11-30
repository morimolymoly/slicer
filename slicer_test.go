package slicer

import "testing"

func Test0SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(0, 100)
	if len(index) != 0 {
		t.Error("0 Case failed!")
	}
}
func Test1SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(1, 100)
	if len(index) != 1 {
		t.Error("1 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 1 {
		t.Error("1 Case failed!")
	}
}

func Test98SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(98, 100)
	if len(index) != 1 {
		t.Error("98 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 98 {
		t.Error("98 Case failed!")
	}
}

func Test99SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(99, 100)
	if len(index) != 1 {
		t.Error("99 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 99 {
		t.Error("99 Case failed!")
	}
}

func Test100SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(100, 100)
	if len(index) != 1 {
		t.Error("100 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 100 {
		t.Error("100 Case failed!")
	}
}

func Test101SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(101, 100)
	if len(index) != 2 {
		t.Error("101 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 100 {
		t.Error("101 Case failed!")
	}
	if index[1].Start != 100 || index[1].End != 101 {
		t.Error("101 Case failed!")
	}
}

func Test102SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(102, 100)
	if len(index) != 2 {
		t.Error("102 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 100 {
		t.Error("102 Case failed!")
	}
	if index[1].Start != 100 || index[1].End != 102 {
		t.Error("102 Case failed!")
	}
}

func Test200SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(200, 100)
	if len(index) != 2 {
		t.Error("200 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 100 {
		t.Error("200 Case failed!")
	}
	if index[1].Start != 100 || index[1].End != 200 {
		t.Error("200 Case failed!")
	}
}

func Test201SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(201, 100)
	if len(index) != 3 {
		t.Error("201 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 100 {
		t.Error("201 Case failed!")
	}
	if index[1].Start != 100 || index[1].End != 200 {
		t.Error("201 Case failed!")
	}
	if index[2].Start != 200 || index[2].End != 201 {
		t.Error("201 Case failed!")
	}
}

func Test210SliceAnd100Buffer(t *testing.T) {
	index := GetSliceIndex(210, 100)
	if len(index) != 3 {
		t.Error("210 Case failed!")
	}
	if index[0].Start != 0 || index[0].End != 100 {
		t.Error("210 Case failed!")
	}
	if index[1].Start != 100 || index[1].End != 200 {
		t.Error("210 Case failed!")
	}
	if index[2].Start != 200 || index[2].End != 210 {
		t.Error("210 Case failed!")
	}
}
