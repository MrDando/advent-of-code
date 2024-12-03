# Import because default Python regex engine doesn't support variable length patterns in lookbehinds
import regex as re

file = open("input.txt", "r")

computer_memory = file.read()
pattern_match = re.findall(r"(?<!(?:don't\(\))(?:[^d]||d(?!o\(\)))+?)mul\((\d+),(\d+)\)", computer_memory) # FML

result = 0
for touple in pattern_match:
    product = int(touple[0]) * int(touple[1])
    result += product

print(result)