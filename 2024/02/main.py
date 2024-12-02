with open("in.1") as f:
    input = f.read().strip()
    lines = input.split("\n")

p1 = 0
p2 = 0

def safe_report(levels, negative):
    for i in range(1, len(levels)):
        a, b = levels[i], levels[i-1]
        diff = abs(a - b)
        if negative:
            diff = b - a
        else:
            diff = a - b
        if diff <= 0 or diff > 3:
            return (False, i-1, i)
    return (True, -1, -1)


def check_positive(levels):
    for i in range(1, len(levels)):
        if levels[i] < levels[i-1]:
            return (False, i-1, i)
    return (True, -1, -1)


def check_negative(levels):
    for i in range(1, len(levels)):
        if levels[i] > levels[i-1]:
            return (False, i-1, i)
    return (True, -1, -1)


for line in lines:
    levels = list(map(int, line.split()))

    positive, _, _ = check_positive(levels)
    negative, _, _ = check_negative(levels)

    if not positive and not negative:
        continue

    pos_safe, _, _ = safe_report(levels, False)
    neg_safe, _, _ = safe_report(levels, True)
    if pos_safe or neg_safe:
        p1 += 1

for line in lines:
    levels = list(map(int, line.split()))

    positive, badp_i, badp_j = check_positive(levels)
    negative, badn_i, badn_j = check_negative(levels)

    safe = False
    if positive:
        safe, bad_i, bad_j = safe_report(levels, False)
        if safe:
            p2 += 1
        else:
            ilevels = levels.copy()
            jlevels = levels.copy()
            del ilevels[bad_i]
            del jlevels[bad_j]

            isafe, _, _ = safe_report(ilevels, False)
            jsafe, _, _ = safe_report(jlevels, False)

            if isafe or jsafe:
                p2 += 1
    elif negative:
        safe, bad_i, bad_j = safe_report(levels, True)
        if safe:
            p2 += 1
        else:
            ilevels = levels.copy()
            jlevels = levels.copy()
            del ilevels[bad_i]
            del jlevels[bad_j]

            isafe, _, _ = safe_report(ilevels, True)
            jsafe, _, _ = safe_report(jlevels, True)

            if isafe or jsafe:
                p2 += 1
    else:
        pilevels = levels.copy()
        pjlevels = levels.copy()
        nilevels = levels.copy()
        njlevels = levels.copy()
        del pilevels[badp_i]
        del pjlevels[badp_j]
        del nilevels[badn_i]
        del njlevels[badn_j]

        positive_i, _, _ = check_positive(pilevels)
        positive_j, _, _ = check_positive(pjlevels)

        negative_i, _, _ = check_negative(nilevels)
        negative_j, _, _ = check_negative(njlevels)

        if positive_i:
            safe, _, _ = safe_report(pilevels, False)
            if safe:
                p2 += 1
                continue
        if positive_j:
            safe, _, _ = safe_report(pjlevels, False)
            if safe:
                p2 += 1
                continue
        if negative_i:
            safe, _, _ = safe_report(nilevels, True)
            if safe:
                p2 += 1
                continue
        if negative_j:
            safe, _, _ = safe_report(njlevels, True)
            if safe:
                p2 += 1
                continue

print(p1)
print(p2)
