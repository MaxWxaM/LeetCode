func climbStairs(n int) int {
    
    if n == 1{
        return 1
    }else if n == 2{
        return 2
    }else{
        result, prev, prevTwo  := 0, 2, 1
        
        for i := 3; i <= n; i++ {
            result = prev + prevTwo
            prevTwo = prev
            prev = result
        }
        return result
    }
}
