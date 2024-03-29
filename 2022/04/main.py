def run():
    with open("in.1") as f:
        input = f.read()

    subset_count = 0
    overlap_count = 0
    for line in input.split("\n"):
        section_1, section_2 = [list(map(int, x.split('-'))) for x in line.split(",")]
        range_1, range_2 = set(range(section_1[0], section_1[1]+1)), set(range(section_2[0], section_2[1]+1))
        if not range_1.isdisjoint(range_2):
            overlap_count += 1
        if range_1.issubset(range_2) or range_1.issuperset(range_2):
            subset_count += 1

    print(subset_count)

    print(overlap_count)

if __name__ == "__main__":
    run()
