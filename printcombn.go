package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
    for i := 0; i <= 10-n; i++ {
        if n == 1 {
            z01.PrintRune(rune(i + 48))
            if i < 9 {
                z01.PrintRune(',')
                z01.PrintRune(' ')
            } else {
                z01.PrintRune('\n')
            }
        } else {
            Solve(n-1, []int{i})
        }
    }
}

func Solve(prev int, s []int) {
    for i := s[len(s)-1] + 1; i <= 10-prev; i++ {
        if prev == 1 {
            PrintNum(append(s, i))
        } else {
            Solve(prev-1, append(s, i))
        }
    }
}

func PrintNum(nums []int) {
    for i := 0; i < len(nums); i++ {
        z01.PrintRune(rune(nums[i] + 48))
    }
    if nums[0] < 10-len(nums) {
        z01.PrintRune(',')
        z01.PrintRune(' ')
    } else {
        z01.PrintRune('\n')
    }
}
