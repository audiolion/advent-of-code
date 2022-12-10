with open(0) as f:
    input = f.read()
    lines = input.split("\n")

def draw(image, register, cycleCount):
    sprite_pos = [register + n + 40 * (len(image)-1) for n in [-1, 0, 1]]
    if cycleCount-1 in sprite_pos:
        image[-1].append("#")
    else:
        image[-1].append(".")

signal_cycles = [20, 60, 100, 140, 180, 220]
register = 1
cycleCount = 0
strengths = []
image = [[]]

for i, line in enumerate(lines):
    for _ in range(1 if line == "noop" else 2):
        cycleCount += 1
        if cycleCount in signal_cycles:
            strengths.append(register)
        draw(image, register, cycleCount)
        if cycleCount % 40 == 0 and i != len(lines)-1:
            image.append([])

    if line != "noop":
        n = int(line.split(" ")[1])
        register += n

print(sum([x * signal_cycles[i] for i, x in enumerate(strengths)]))
for row in image:
    print("".join(row))
