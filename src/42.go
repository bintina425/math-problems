def square_root(x):
    if x < 0:
        raise ValueError("Cannot compute the square root of a negative number.")
    else:
        y = x ** 0.5
        return y

try:
    print(square_root(4))
except ValueError as e:
    print(e)
