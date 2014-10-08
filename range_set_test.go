package range_set

import (
    "testing"
    "fmt"
)


func TestAdd(t *testing.T) {
    r := RangeSet{}
    if len(r.Ranges) != 0 {
        t.Errorf("Length of Ranges is not 0")
    }

    // nums := []int64{1, 2, 3, 6, 7, 8}
    nums := []int64{1}
    r.AddInts(nums)
    fmt.Printf("r.Ranges: %s", r.Ranges)
}
