from os import path


basepath = path.dirname(__file__)

def run():
    with open(path.abspath(path.join(basepath, "in.1"))) as f:
        input = f.read()

    inventories = input.split("\n\n")
    maxes = [0, 0, 0]
    for inventory in inventories:
        items = [int(item) for item in inventory.split("\n")]
        total_calories = sum(items)

        for i, max in enumerate(maxes):
            if total_calories > max:
                tmp, maxes[i] = maxes[i], total_calories
                for j in range(i+1, len(maxes)):
                    if tmp > maxes[j]:
                        tmp, maxes[j] = maxes[j], tmp
                break


    print(maxes[0])

    print(sum(maxes))


if __name__ == "__main__":
    run()
