
def visible_and_scenic_score(m, i, j):
    v_top, v_left, v_bot, v_right = True, True, True, True
    s_top, s_left, s_bot, s_right = 0, 0, 0, 0
    tree = m[i][j]

    p = i-1
    while p >= 0:
        s_top += 1
        if m[p][j] >= tree:
            v_top = False
            break
        p -= 1

    p = j-1
    while p >= 0:
        s_left += 1
        if m[i][p] >= tree:
            v_left = False
            break
        p -= 1

    p = i+1
    while p < len(m):
        s_bot += 1
        if m[p][j] >= tree:
            v_bot = False
            break
        p += 1

    p = j+1
    while p < len(m[0]):
        s_right += 1
        if m[i][p] >= tree:
            v_right = False
            break
        p += 1

    return (v_top or v_left or v_bot or v_right, s_top * s_left * s_bot * s_right)

with open(0) as f:
    input = f.read()

m = []
for i, line in enumerate(input.split("\n")):
    m.append([])
    for tree in line:
        m[i].append(tree)

visible = 0
score = 0
for i in range(len(m)):
    for j in range(len(m[0])):
        if i-1 < 0 or j-1 < 0 or i+1 >= len(m) or j+1 >= len(m[0]):
            visible += 1
            continue
        local_visible, local_score = visible_and_scenic_score(m, i, j)
        if local_visible:
            visible += 1
        if local_score > score:
            score = local_score

print(visible)
print(score)
