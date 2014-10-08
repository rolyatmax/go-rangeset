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

    r.AddInts([]int64{1, 2, 3})
    fmt.Printf("r.Ranges: %d \n", r.Ranges)

    if r.Ranges[0].Low != 1 || r.Ranges[0].High != 3 {
        t.Errorf("Expected 1-3")
    }

    r.AddInts([]int64{6, 7, 8})
    fmt.Printf("added 6, 7, 8: %d \n", r.Ranges)

    if r.Ranges[1].Low != 6 || r.Ranges[1].High != 8 {
        t.Errorf("Expected 1-3, 6-8")
    }

    r.AddInts([]int64{12, 14, 15, 16, 17})
    fmt.Printf("added 12, 14-17: %d \n", r.Ranges)

    if r.Ranges[2].Low != 12 || r.Ranges[2].High != 12 {
        t.Errorf("Expected 1-3, 6-8, 12-12, 14-17")
    }

    r.AddInts([]int64{4, 5})
    fmt.Printf("added 4, 5: %d \n", r.Ranges)

    if r.Ranges[0].Low != 1 || r.Ranges[0].High != 8 {
        t.Errorf("Expected 1-8, 12-12, 14-17")
    }

    r.AddInts([]int64{13})
    fmt.Printf("added 4, 5: %d \n", r.Ranges)

    if r.Ranges[1].Low != 12 || r.Ranges[1].High != 17 {
        t.Errorf("Expected 1-8, 12-17")
    }
}
