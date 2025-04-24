def is_perfect_square(n):
    """
    Check if n is a perfect square.
    
    Parameters:
    n (int): The number to check
    
    Returns:
    bool: True if n is a perfect square, False otherwise.
    """
    root = int(n ** 0.5)
    return root * root == n

def solve_problem():
    """
    Solve a math problem for the given solution to math-problems (Solutions to math problems for my math class).
    
    Returns:
    str: The solution to the problem.
    """
    # Example math problem
    problem = "A mathematician is trying to prove that every square number is congruent to 0 or 1 modulo 4."
    return f"The solution to the mathematical problem '{problem}' is {is_perfect_square(problem)}."

# Uncomment to test the function
# print(solve_problem())
