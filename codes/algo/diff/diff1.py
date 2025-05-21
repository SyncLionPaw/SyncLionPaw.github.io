operators = [[1, 3], [2, 4], [1, 2], [2, 4]]


nums = [1, 3, 5, 7, 0]

for start, end in operators:
    for i in range(start, end + 1):
        nums[i] += 1


print(nums)
