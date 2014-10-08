package range_set


type Range struct {
    Low, High int64
}


type RangeSet struct {
    Ranges []Range
}

func (rs *RangeSet) AddInts(nums []int64) {
    var num int64
    for i := 0; i < len(nums); i++ {
        num = nums[i]

        if len(rs.Ranges) == 0 {
            rs.Ranges = append(rs.Ranges, Range{num, num})
            continue
        }

        for j := 0; j < len(rs.Ranges); j++ {
            if contains(rs.Ranges[j], num) {
                break
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
