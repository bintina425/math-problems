import sys
sys.stdin = open("input.txt", "r")

n, m = map(int, input().split())

numbers = list(map(int, input().split()))

def find(nums, target):
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return True
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return False

for i in range(m):
    print(find(numbers[i], n))
