max_red = 12
max_green = 13
max_blue = 14

p1 = 0
p2 = 0

with open(0) as f:
    input = f.read()
    lines = input.split("\n")

ids = []
powers = []
for line in lines:
    game, record = line.split(":")
    id = game.split("Game ")[1]
    pulls = record.split(";")
    reds, greens, blues = [], [], []
    valid = True
    for pull in pulls:
        cubes = pull.split(", ")
        for cube in cubes:
            if "blue" in cube:
                blue = int(cube.split(" blue")[0])
                blues.append(blue)
                if blue > max_blue:
                    valid = False
            if "green" in cube:
                green = int(cube.split(" green")[0])
                greens.append(green)
                if green > max_green:
                    valid = False
            if "red" in cube:
                red = int(cube.split(" red")[0])
                reds.append(red)
                if red > max_red:
                    valid = False

    if valid:
        ids.append(int(id))

    min_red = max(reds)
    min_green = max(greens)
    min_blue = max(blues)
    power = min_red * min_green * min_blue
    powers.append(power)


p1 = sum(ids)
p2 = sum(powers)

print(p1)
print(p2)
