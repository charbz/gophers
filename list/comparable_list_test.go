package list

import (
	"slices"
	"testing"
)

func TestComparableList_Contains(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		value int
		want  bool
	}{
		{
			name:  "contains value",
			slice: []int{1, 2, 3, 4, 5},
			value: 3,
			want:  true,
		},
		{
			name:  "does not contain value",
			slice: []int{1, 2, 3, 4, 5},
			value: 6,
			want:  false,
		},
		{
			name:  "empty list",
			slice: []int{},
			value: 1,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewComparableList(tt.slice)
			got := l.Contains(tt.value)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparableList_Distinct(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "no duplicates",
			slice: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "with duplicates",
			slice: []int{1, 2, 2, 3, 3, 3, 4, 5, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "empty list",
			slice: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewComparableList(tt.slice)
			got := l.Distinct()
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Distinct() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestComparableList_Diff(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   []int
	}{
		{
			name:   "different lists",
			slice1: []int{1, 2, 3, 4, 5},
			slice2: []int{4, 5, 6, 7},
			want:   []int{1, 2, 3},
		},
		{
			name:   "identical lists",
			slice1: []int{1, 2, 3},
			slice2: []int{1, 2, 3},
			want:   []int{},
		},
		{
			name:   "empty first list",
			slice1: []int{},
			slice2: []int{1, 2, 3},
			want:   []int{},
		},
		{
			name:   "empty second list",
			slice1: []int{1, 2, 3},
			slice2: []int{},
			want:   []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := NewComparableList(tt.slice1)
			l2 := NewComparableList(tt.slice2)
			got := l1.Diff(l2)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Diff() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestComparableList_Equals(t *testing.T) {
	tests := []struct {
		name   string
		slice1 []int
		slice2 []int
		want   bool
	}{
		{
			name:   "equal lists",
			slice1: []int{1, 2, 3},
			slice2: []int{1, 2, 3},
			want:   true,
		},
		{
			name:   "different values",
			slice1: []int{1, 2, 3},
			slice2: []int{1, 2, 4},
			want:   false,
		},
		{
			name:   "different lengths",
			slice1: []int{1, 2, 3},
			slice2: []int{1, 2},
			want:   false,
		},
		{
			name:   "empty lists",
			slice1: []int{},
			slice2: []int{},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := NewComparableList(tt.slice1)
			l2 := NewComparableList(tt.slice2)
			got := l1.Equals(l2)
			if got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparableList_IndexOf(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		value int
		want  int
	}{
		{
			name:  "value exists",
			slice: []int{1, 2, 3, 4, 5},
			value: 3,
			want:  2,
		},
		{
			name:  "value does not exist",
			slice: []int{1, 2, 3, 4, 5},
			value: 6,
			want:  -1,
		},
		{
			name:  "empty list",
			slice: []int{},
			value: 1,
			want:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewComparableList(tt.slice)
			got := l.IndexOf(tt.value)
			if got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparableList_LastIndexOf(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		value int
		want  int
	}{
		{
			name:  "value exists once",
			slice: []int{1, 2, 3, 4, 5},
			value: 3,
			want:  2,
		},
		{
			name:  "value exists multiple times",
			slice: []int{1, 2, 3, 2, 4},
			value: 2,
			want:  3,
		},
		{
			name:  "value does not exist",
			slice: []int{1, 2, 3, 4, 5},
			value: 6,
			want:  -1,
		},
		{
			name:  "empty list",
			slice: []int{},
			value: 1,
			want:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewComparableList(tt.slice)
			got := l.LastIndexOf(tt.value)
			if got != tt.want {
				t.Errorf("LastIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparableList_Max(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		want    int
		wantErr bool
	}{
		{
			name:    "non-empty list",
			slice:   []int{1, 5, 3, 4, 2},
			want:    5,
			wantErr: false,
		},
		{
			name:    "single element",
			slice:   []int{1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "empty list",
			slice:   []int{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewComparableList(tt.slice)
			got, err := l.Max()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Max() = %v, want error", got)
				}
			} else {
				if err != nil {
					t.Errorf("Max() = %v, want no error", got)
				}
				if got != tt.want {
					t.Errorf("Max() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestComparableList_Min(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		want    int
		wantErr bool
	}{
		{
			name:    "non-empty list",
			slice:   []int{5, 1, 3, 4, 2},
			want:    1,
			wantErr: false,
		},
		{
			name:    "single element",
			slice:   []int{1},
			want:    1,
			wantErr: false,
		},
		{
			name:    "empty list",
			slice:   []int{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewComparableList(tt.slice)
			got, err := l.Min()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Min() = %v, want error", got)
				}
			} else {
				if err != nil {
					t.Errorf("Min() = %v, want no error", got)
				}
				if got != tt.want {
					t.Errorf("Min() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestComparableList_Sum(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{
			name:  "non-empty list",
			slice: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "single element",
			slice: []int{5},
			want:  5,
		},
		{
			name:  "empty list",
			slice: []int{},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewComparableList(tt.slice)
			got := l.Sum()
			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComparableList_StartsWith(t *testing.T) {
	tests := []struct {
		name       string
		list1      []int
		list2      []int
		startsWith bool
	}{
		{
			name:       "starts with matching elements",
			list1:      []int{1, 2, 3, 4},
			list2:      []int{1, 2},
			startsWith: true,
		},
		{
			name:       "does not start with different elements",
			list1:      []int{1, 2, 3, 4},
			list2:      []int{2, 3},
			startsWith: false,
		},
		{
			name:       "empty list2 (always true)",
			list1:      []int{1, 2, 3, 4},
			list2:      []int{},
			startsWith: true,
		},
		{
			name:       "list1 shorter than list2",
			list1:      []int{1, 2},
			list2:      []int{1, 2, 3},
			startsWith: false,
		},
		{
			name:       "both lists empty",
			list1:      []int{},
			list2:      []int{},
			startsWith: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := NewComparableList(tt.list1)
			l2 := NewComparableList(tt.list2)
			if got := l1.StartsWith(l2); got != tt.startsWith {
				t.Errorf("StartsWith() = %v, want %v", got, tt.startsWith)
			}
		})
	}
}

func TestComparableList_EndsWith(t *testing.T) {
	tests := []struct {
		name      string
		list1     []int
		list2     []int
		endsWith  bool
	}{
		{
			name:      "ends with matching elements",
			list1:     []int{1, 2, 3, 4},
			list2:     []int{3, 4},
			endsWith:  true,
		},
		{
			name:      "does not end with different elements",
			list1:     []int{1, 2, 3, 4},
			list2:     []int{2, 3},
			endsWith:  false,
		},
		{
			name:      "empty list2 (always true)",
			list1:     []int{1, 2, 3, 4},
			list2:     []int{},
			endsWith:  true,
		},
		{
			name:      "list1 shorter than list2",
			list1:     []int{3, 4},
			list2:     []int{2, 3, 4},
			endsWith:  false,
		},
		{
			name:      "both lists empty",
			list1:     []int{},
			list2:     []int{},
			endsWith:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := NewComparableList(tt.list1)
			l2 := NewComparableList(tt.list2)
			if got := l1.EndsWith(l2); got != tt.endsWith {
				t.Errorf("EndsWith() = %v, want %v", got, tt.endsWith)
			}
		})
	}
}