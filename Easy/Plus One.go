// 1-st solution
// Time complexity - O(n)
// Space complexity - O(n)

func plusOne(digits []int) []int {
    n := len(digits)
    
    for i := n - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }
        
        digits[i] = 0
    }

    return append([]int{1}, digits...)
}

// 2-nd solution
// Time complexity - O(n)
// Space complexity - O(n)

func plusOne(digits []int) []int {
    n := len(digits)

    for i := n - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++
            return digits
        }

        digits[i] = 0
    }

    result := make([]int, n+1)
    result[0] = 1

    copy(result[1:], digits)

    return result
}
