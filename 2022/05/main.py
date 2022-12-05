def print_stacks(stacks):
    for i, stack in enumerate(stacks):
        print(i, stack)


def parse_message(stacks):
    message = ""
    for stack in stacks:
        if len(stack) > 0:
            message += stack[-1]
        else:
            message += " "

    return message.replace("[", "").replace("]", "")


def parse_stacks(stack_data):
    stack_layout = stack_data.split("\n")
    stack_layout.reverse()
    stacks = []
    for i, line in enumerate(stack_layout):
        if i == 0:
            copy = line.strip()
            start = int(copy[0])
            end = int(copy[-1])
            for _ in range(start-1, end):
                stacks.append([])
            continue

        imputed_len = len(stacks) * 3 + len(stacks) - 1
        imputed = f"{line}{(imputed_len - len(line)) * ' '}"
        i = 0
        while i < len(imputed):
            if imputed[i:i+3] == "   ":
                imputed = f"{imputed[0:i]}[_]{imputed[i+3:len(imputed)]}"
            i += 4

        crates = imputed.split(" ")
        for i, crate in enumerate(crates):
            stacks[i].append(crate)

    for i, stack in enumerate(stacks):
        stacks[i] = [s for s in stack if s != "[_]"]

    return stacks


def parse_instruction(instruction):
    parts = instruction.replace("move ", "").replace("from ", "").replace("to ", "").split(" ")
    num_crates, from_crate_idx, to_crate_idx = int(parts[0]), int(parts[1])-1, int(parts[2])-1
    return num_crates, from_crate_idx, to_crate_idx

def run():
    with open("in.1") as f:
        input = f.read()

    stack_data, instructions = input.split("\n\n")

    stacks_1 = parse_stacks(stack_data)
    for instruction in instructions.split("\n"):
        num_crates, from_crate_idx, to_crate_idx = parse_instruction(instruction)
        from_crate = stacks_1[from_crate_idx]
        to_crate = stacks_1[to_crate_idx]
        for _ in range(0, num_crates):
            if len(from_crate) > 0:
                crate = from_crate.pop()
                to_crate.append(crate)

    print(parse_message(stacks_1))

    stacks_2 = parse_stacks(stack_data)
    for instruction in instructions.split("\n"):
        num_crates, from_crate_idx, to_crate_idx = parse_instruction(instruction)
        from_crate = stacks_2[from_crate_idx]
        to_crate = stacks_2[to_crate_idx]
        to_crate.extend(from_crate[-num_crates:])
        stacks_2[from_crate_idx] = from_crate[:-num_crates]

    print(parse_message(stacks_2))

if __name__ == "__main__":
    run()
