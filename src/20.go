def add(x, y):
    while True:
        if x + 1 == y or x - 1 == y or (x * 2) % 3 == 0 or (y - x - 1) % 3 == 0:
            return x + 1
        else:
            return x - 1

