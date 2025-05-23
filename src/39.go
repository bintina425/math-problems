def is_prime(num):
    if num <= 1:
        return False
    for i in range(2, int(num ** 0.5) + 1):
        if num % i == 0:
            return False
    return True

def square_root(n):
    import math
    return math.sqrt(n)

# Example usage:
num = 7
is_prime(num)
square_root(9)
