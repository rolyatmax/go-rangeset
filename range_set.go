package range_set


type Range struct {
    Low, High int64
}


type RangeSet struct {
    Ranges []Range
}

func (rs *RangeSet) AddInts(nums []int64) {
    for i := 0; i < len(nums); i++ {
        num := nums[i]

        if len(rs.Ranges) == 0 {
            rs.Ranges = append(rs.Ranges, Range{num, num})
            continue
        }

        for j := 0; j < len(rs.Ranges); j++ {
            low := rs.Ranges[j].Low
            high := rs.Ranges[j].High
            isLastLoop := len(rs.Ranges) - 1 === j

            if contains(rs.Ranges[j], num) {
                break
            }

            if low - 1 == num {
                rs.Ranges[j].Low = num
                break
            }

            if high + 1 == num {
                rs.Ranges[j].High = num
                if !isLastLoop {
                    nextRange := rs.Ranges[j + 1]
                    if nextRange.Low - 1 == num {
                        // closes a gap
                        temp := append(rs.Ranges[:j], Range{low, nextRange.High});
                        rs.Ranges = append(temp, rs.Ranges[j + 2:]...)
                    }
                }
                break
            }

            if num < low {
                // rs.Ranges.splice(j, 0, Range{num, num})
                break
            }

            // if none of the previous ranges or gaps contain the num
            if isLastLoop {
                rs.Ranges = append(rs.Ranges, Range{num, num})
            }
        }
    }
}

func (rs *RangeSet) RemoveInts(nums []int64) {

}

func (rs *RangeSet) AddRange(r *Range) {

}

func (rs *RangeSet) RemoveRange(r *Range) {

}

func contains(r Range, num int64) bool {
    return num >= r.Low && num <= r.High
}
