import re


def parse_stacks(stack_data):
    stacks = []
    lines = stack_data.split("\n")
    for line in lines:
        stacks.append([line[i] if len(line) >= i else ' ' for i in range(1, len(lines[-1]), 4)])
    stacks = zip(*stacks[:-1][::-1])
    stacks = [list("".join(s).strip()) for s in stacks]
    return stacks


def run():
    with open("in.1") as f:
        input = f.read()

    stack_data, instructions = input.split("\n\n")

    s1 = parse_stacks(stack_data)
    for instruction in instructions.split("\n"):
        n, src, dest = map(int, re.findall("\\d+", instruction))
        for _ in range(0, n):
            if len(s1[src-1]) > 0:
                s1[dest-1].append(s1[src-1].pop())

    print("".join([x[-1] for x in s1]))

    s2 = parse_stacks(stack_data)
    for instruction in instructions.split("\n"):
        n, src, dest = map(int, re.findall("\\d+", instruction))
        s2[dest-1].extend(s2[src-1][-n:])
        s2[src-1] = s2[src-1][:-n]

    print("".join([x[-1] for x in s2]))


if __name__ == "__main__":
    run()
