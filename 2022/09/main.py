with open("in.1") as f:
    input = f.read()

DIAG = [(+1, +1), (+1, -1), (-1, -1), (-1, +1)]
def is_diag(tail, head):
    for r, c in DIAG:
        if (tail[0] + r, tail[1] + c) == head:
            return True
    return False

ADJ = [(+1, 0), (-1, 0), (0, +1), (0, -1)]
def is_adj(tail, head):
    for r, c in ADJ:
        if (tail[0] + r, tail[1] + c) == head:
            return True
    return False

def is_overlap(tail, head):
    return tail == head

def is_touching(tail, head):
    return is_diag(tail, head) or is_adj(tail, head) or is_overlap(tail, head)

def move_required(tail, head):
    if tail[0] == head[0] or tail[1] == head[1]:
        for r, c in ADJ:
            if (tail[0] + r, tail[1] + c) == (head[0] - r, head[1] - c):
                return (r, c)

    for r, c in DIAG:
        if is_touching((tail[0] + r, tail[1] + c), head):
            return (r, c)

def run(num_knots):
    visited = set((0, 0))

    def update_knots(knots):
        for k in range(len(knots)-1):
            if is_touching(knots[k+1], knots[k]):
                return
            r, c = move_required(knots[k+1], knots[k])
            knots[k+1] = (knots[k+1][0] + r, knots[k+1][1] + c)
            if k+1 == len(knots) - 1:
                visited.add(knots[k+1])

    knots = [(0, 0) for _ in range(num_knots)]
    for line in input.split("\n"):
        d, n = line.split(" ")
        n = int(n)
        if d == "U":
            for _ in range(n):
                knots[0] = (knots[0][0] + 1, knots[0][1])
                update_knots(knots)
        elif d == "R":
            for _ in range(n):
                knots[0] = (knots[0][0], knots[0][1] + 1)
                update_knots(knots)
        elif d == "D":
            for _ in range(n):
                knots[0] = (knots[0][0] - 1, knots[0][1])
                update_knots(knots)
        elif d == "L":
            for _ in range(n):
                knots[0] = (knots[0][0], knots[0][1] - 1)
                update_knots(knots)

    print(len(visited))

run(2)
run(10)
