

with open(0) as f:
    input = f.read()
    lines = input.split("\n")

p1 = 0
p2 = 0

parts = []
gears = []
coords = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]

V = {}
for i, line in enumerate(lines):
    part = ""
    ok = False
    for j, ch in enumerate(line):
        if not ch.isdigit():
            if ok:
                for k in range(len(part)):
                    V[(i, j-k-1)] = int(part)
                parts.append(int(part))
            part = ""
            ok = False
            continue

        part += ch
        for x, y in coords:
            if i+x >= 0 and i+x < len(lines):
                if j+y >= 0 and j+y < len(lines[i+x]):
                    if not lines[i+x][j+y].isdigit() and not lines[i+x][j+y] == ".":
                        ok = True
    if ok:
        for k in range(len(part)):
            V[(i, len(line)-k-1)] = int(part)
        parts.append(int(part))
    part = ""
    ok = False

for i, line in enumerate(lines):
    for j, ch in enumerate(line):
        if ch == "*":
            S = set()
            for x, y in coords:
                coord = (i+x, j+y)
                if V.get(coord, None) is not None:
                    S.add(V.get(coord))
            S = list(S)
            if len(S) == 2:
                gears.append(S[0] * S[1])


p1 = sum(parts)
p2 = sum(gears)
print(p1)
print(p2)
