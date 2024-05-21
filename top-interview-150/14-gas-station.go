package leetcode

import "fmt"

func ccr(gas []int, cost []int, start int) int {
    rg := gas[start]
    l := len(gas)
    var idx int

    var counter int

    for i := 0; i < l; i++ {
        counter++
        idx = (start + i) % l
        var next int
        if idx == l - 1 {
            next = 0
        } else {
            next = idx + 1
        }
        rg = rg - cost[idx]
        
        if rg < 0 { return -1 }
        
        if counter < l {
            rg = rg + gas[next]
        }
    }

    return start
}
func canCompleteCircuit(gas []int, cost []int) int {
    for idx, g := range gas {
        if g > 0 && g >= cost[idx] {
            start := idx
            result := ccr(gas, cost, start)
            if result != -1 { return result }
        }
    }

    return -1
}
