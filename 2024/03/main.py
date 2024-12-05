import re

with open("in.1") as f:
    input = "do()"+f.read().strip()
    input = "".join(input.split("\n"))

p1 = 0
p2 = 0

matches = re.findall(r'mul\([0-9]{1,},[0-9]{1,}\)', input)
for match in matches:
    match = match[4:len(match)-1]
    left, right = match.split(',')
    p1 += int(left)*int(right)

matches = re.findall(r'do\(\).+?(?=(?=don\'t\(\))|(?=$))', input)

for match in matches:
    matches2 = re.findall(r'mul\([0-9]{1,},[0-9]{1,}\)', match)
    for m in matches2:
        m = m[4:len(m)-1]
        left, right = m.split(',')
        p2 += int(left)*int(right)

print(p1)
print(p2)
