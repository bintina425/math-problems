def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False
    return True

# Generate random prime numbers
primes = [i for i in range(100) if is_prime(i)]
print(primes)
