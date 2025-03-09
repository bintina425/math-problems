func RandomMathProblem() int {
    // Generate a random number between 1 and 10
    x := rand.Intn(10) + 1
    // Generate a random operation (+, -, *, /)
    op := rand.Intn(4)
    switch op {
        case 0:
            return x * 2
        case 1:
            return x + 3
        case 2:
            return x - 5
        default:
            return x / 7
    }
}
