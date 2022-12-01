from os import path
import heapq

basepath = path.dirname(__file__)

def run():
    with open(path.abspath(path.join(basepath, "in.1"))) as f:
        input = f.read()

    inventories = input.split("\n\n")
    maxes = [-3,-2,-1]
    for inventory in inventories:
        items = [int(item) for item in inventory.split("\n")]
        total_calories = sum(items)
        heapq.heappushpop(maxes, total_calories)

    print(heapq.nlargest(1, maxes)[0])

    print(sum(heapq.nlargest(3, maxes)))


if __name__ == "__main__":
    run()
