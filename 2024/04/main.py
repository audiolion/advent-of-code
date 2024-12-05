from collections import defaultdict

with open(0) as f:
    input = f.read().strip()
    lines = input.split("\n")

p1 = 0
p2 = 0

blank = list('Z'*(len(lines[0])+6))

G = [blank, blank, blank]
for line in lines:
    G.append(list('ZZZ'+line+'ZZZ'))

G.append(blank)
G.append(blank)
G.append(blank)

for i, row in enumerate(G):
    for j, c in enumerate(row):
        if c == 'X':
            # backwards
            if G[i-1][j] == 'M' and G[i-2][j] == 'A' and G[i-3][j] == 'S':
                p1 += 1
            # forwards
            if G[i+1][j] == 'M' and G[i+2][j] == 'A' and G[i+3][j] == 'S':
                p1 += 1
            # up
            if G[i][j-1] == 'M' and G[i][j-2] == 'A' and G[i][j-3] == 'S':
                p1 += 1
            # down
            if G[i][j+1] == 'M' and G[i][j+2] == 'A' and G[i][j+3] == 'S':
                p1 += 1
            # left up diag
            if G[i-1][j-1] == 'M' and G[i-2][j-2] == 'A' and G[i-3][j-3] == 'S':
                p1 += 1
            # left down diag
            if G[i-1][j+1] == 'M' and G[i-2][j+2] == 'A' and G[i-3][j+3] == 'S':
                p1 += 1
            # right down diag
            if G[i+1][j+1] == 'M' and G[i+2][j+2] == 'A' and G[i+3][j+3] == 'S':
                p1 += 1
            # right up diag
            if G[i+1][j-1] == 'M' and G[i+2][j-2] == 'A' and G[i+3][j-3] == 'S':
                p1 += 1

for i, row in enumerate(G):
    for j, c in enumerate(row):
        if c == 'A':
            V = defaultdict(int)
            V[G[i-1][j-1]] += 1
            V[G[i-1][j+1]] += 1
            V[G[i+1][j-1]] += 1
            V[G[i+1][j+1]] += 1
            if V['S'] == 2 and V['M'] == 2 and G[i-1][j-1] != G[i+1][j+1]:
                p2 += 1

print(p1)
print(p2)
