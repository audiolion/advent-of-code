from functools import reduce


def run():
    with open("in.1", "r") as f:
        input = f.read()

    lines = input.split("\n")

    # Priority -> a-z = 1..26, A-Z = 27..52
    def get_priority(item):
        if ord(item) >= ord('a'):
            return ord(item) - ord('a') + 1
        else:
            return ord(item) - ord('A') + 26 + 1

    priority = 0
    for line in lines:
        mid = int(len(line) / 2)
        compartment_1 = line[:mid]
        compartment_2 = line[mid:]
        s1 = set(compartment_1)
        s2 = set(compartment_2)
        common_item = s1.intersection(s2).pop()
        priority += get_priority(common_item)

    print(priority)


    group_size = 3
    group = [None for i in range(0, group_size)]

    priority_2 = 0
    for i, line in enumerate(lines):
        group[i%group_size] = set(line)
        next_group = (i+1) % group_size == 0
        if next_group:
            common_item = reduce(lambda a, b: a.intersection(b), group, set(group[0])).pop()
            priority_2 += get_priority(common_item)

    print(priority_2)



if __name__ == "__main__":
    run()
