import re

file = open("input.txt", "r")

computer_memory = file.read()
pattern_match = re.findall(r"mul\((\d+),(\d+)\)", computer_memory)

result = 0
for touple in pattern_match:
    product = int(touple[0]) * int(touple[1])
    result += product

print(result)