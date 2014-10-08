package range_set


type Range struct {
    Low, High int64
}


type RangeSet struct {
    Ranges []Range
}

func (rs *RangeSet) AddInts(nums []int64) {
    for _, num := range nums {
        if len(rs.Ranges) == 0 {
            rs.Ranges = append(rs.Ranges, Range{num, num})
            continue
        }

        for j := 0; j < len(rs.Ranges); j++ {
            low := rs.Ranges[j].Low
            high := rs.Ranges[j].High
            isLastLoop := len(rs.Ranges) - 1 == j

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
                        rs.Ranges = splice(rs.Ranges, j, 2, Range{low, nextRange.High})
                    }
                }
                break
            }

            if num < low {
                rs.Ranges = splice(rs.Ranges, j, 0, Range{num, num})
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
    for _, num := range nums {
        for j := 0; j < len(rs.Ranges); j++ {
            if !contains(rs.Ranges[j], num) {
                continue
            }

            low := rs.Ranges[j].Low
            high := rs.Ranges[j].High

            if low == num && high == num {
                rs.Ranges = remove(rs.Ranges, j, 1)
            } else if low == num {
                rs.Ranges[j].Low = low + 1
            } else if high == num {
                rs.Ranges[j].High = high - 1
            } else {
                rs.Ranges = splice(rs.Ranges, j, 1, Range{low, num - 1})
                rs.Ranges = splice(rs.Ranges, j + 1, 0, Range{num + 1, high})
            }
            break
        }
    }
}

func (rs *RangeSet) AddRange(r *Range) {

}

func (rs *RangeSet) RemoveRange(r *Range) {

}

func contains(r Range, num int64) bool {
    return num >= r.Low && num <= r.High
}

func splice(ranges []Range, startIdx int, elCount int, toInsert Range) []Range {
    temp := append(ranges[:startIdx], toInsert)
    return append(temp, ranges[startIdx + elCount:]...)
}

func remove(ranges []Range, startIdx int, elCount int) []Range {
    return append(ranges[:startIdx], ranges[startIdx + elCount:]...)
}
