package range_set

import (
    "testing"
    "fmt"
)


func TestAddRemoveInts(t *testing.T) {
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


    r.RemoveInts([]int64{14, 15})
    fmt.Printf("removed 14, 15: %d \n", r.Ranges)

    if r.Ranges[3].Low != 16 || r.Ranges[3].High != 17 {
        t.Errorf("Expected 1-3, 6-8, 12-12, 16-17")
    }


    r.AddInts([]int64{4, 5})
    fmt.Printf("added 4, 5: %d \n", r.Ranges)

    if r.Ranges[0].Low != 1 || r.Ranges[0].High != 8 {
        t.Errorf("Expected 1-8, 12-12, 16-17")
    }


    r.AddInts([]int64{13})
    fmt.Printf("added 13: %d \n", r.Ranges)

    if r.Ranges[1].Low != 12 || r.Ranges[1].High != 13 {
        t.Errorf("Expected 1-8, 12-13, 16-17")
    }


    r.RemoveInts([]int64{12, 13, 14, 15, 16, 17})
    fmt.Printf("removed 12-17: %d \n", r.Ranges)

    if len(r.Ranges) != 1 {
        t.Errorf("Expected 1-8")
    }


    r.RemoveInts([]int64{2, 3})
    fmt.Printf("removed 2, 3: %d \n", r.Ranges)

    if len(r.Ranges) != 2 && r.Ranges[0].Low != 1 && r.Ranges[0].High != 1 {
        t.Errorf("Expected 1-1, 4-8")
    }
    if r.Ranges[1].Low != 4 && r.Ranges[1].High != 8 {
        t.Errorf("Expected 1-1, 4-8")
    }
}
