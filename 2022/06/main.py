def run():
    with open(0) as f:
        input = f.read()

    def find_unique_window_idx(input, size):
        j = size
        for i in range(0, len(input)):
            if len(set(input[i:j])) == size:
                return j
            j += 1
        return -1


    for line in input.split("\n"):
        start_of_packet_idx = find_unique_window_idx(line, 4)
        print(start_of_packet_idx)

        start_of_message_idx = find_unique_window_idx(line, 14)
        print(start_of_message_idx)


if __name__ == "__main__":
    run()
