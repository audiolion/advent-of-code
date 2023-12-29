from collections import defaultdict

with open(0) as f:
    input = f.read()
    lines = input.split("\n")

p1 = 0
p2 = 0
V = defaultdict(int)
for line in lines:
    card, rest = line.split(": ")
    card = int(card.split()[1])
    winners, drawn = rest.split(" | ")
    winners = [int(n) for n in winners.split()]
    drawn = [int(n) for n in drawn.split()]
    # print(card, winners, "|", drawn)
    total = 0
    matches = 0
    for n in drawn:
        if n in winners:
            matches += 1
            if total < 2:
                total += 1
            else:
                total = total * 2
    p1 += total


    for k in range(card+1, card+1+matches):
        copies = V[card]
        V[k] += 1 + copies
    V[card] += 1

p2 = sum(V.values())


print(p1)
print(p2)
