from collections import defaultdict

with open(0) as f:
    input = f.read().strip()
    lines = input.split("\n")

p1 = 0
p2 = 0

left = []
right = []
for line in lines:
    l, r = line.split('   ')
    left.append(int(l))
    right.append(int(r))

left.sort()
right.sort()

V = defaultdict(int)
for n in right:
    V[n] += 1

for i in range(len(left)):
    p1 += abs(left[i] - right[i])
    p2 += left[i] * V[left[i]]


def abs(x):
    if x < 0:
        return -x
    return x

print(p1)
print(p2)
