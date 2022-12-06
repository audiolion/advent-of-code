def run():
    with open(0) as f:
        input = f.read()

    for line in input.split("\n"):
        start_of_packet_idx = 0
        j = 4
        for i in range(0, len(line)):
            if len(set(line[i:j])) == 4:
                start_of_packet_idx = j
                break
            j += 1

        print(start_of_packet_idx)

        j = 14
        start_of_message_idx = 0
        for i in range(0, len(line)):
            if len(set(line[i:j])) == 14:
                start_of_message_idx = j
                break
            j += 1

        print(start_of_message_idx)


if __name__ == "__main__":
    run()
