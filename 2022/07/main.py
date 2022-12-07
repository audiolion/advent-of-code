import math


class File:
    size = 0
    name = ""

    def __init__(self, size, name):
        self.size = size
        self.name = name


class Directory:
    def __init__(self, name, parent):
        self.name = name
        self.parent = parent
        self.files = {}
        self.total = 0

    def add_file(self, size, name):
        self.files[name] = File(size, name)
        self.total += int(size)

    def add_dir(self, dir):
        self.files[dir] = Directory(dir, self)


def print_walk(dir, level = 0):
    print(f"{'  ' * level}- {dir.name} (dir)")
    for k, v in dir.files.items():
        if isinstance(v, Directory):
            print_walk(v, level+1)
        else:
            print(f"{'  ' * (level+1)}- {v.name} (file, size={v.size})")


def run():
    with open(0) as f:
        input = f.read()


    fs = None
    curr = None
    is_ls = False
    for line in input.split("\n"):
        if line.startswith("$ cd"):
            is_ls = False
            dir = line.split("$ cd")[1].strip()
            if fs is None:
                fs = Directory(dir, None)
                curr = fs
            elif dir == "..":
                curr = curr.parent
            else:
                curr = curr.files[dir]
        if is_ls:
            if line.startswith("dir"):
                _, dir = line.split(" ")
                curr.add_dir(dir)
            else:
                size, filename = line.split(" ")
                curr.add_file(size, filename)
        if line.startswith("$ ls"):
            is_ls = True

    def dfs(dir, path, totals):
        totals[path] = 0
        for k, v in dir.files.items():
            if isinstance(v, Directory):
                totals[path] += dfs(v, f"{path}{k}/", totals)

        totals[path] += dir.total

        return totals[path]

    totals = {}
    dfs(fs, fs.name, totals)

    max_size = 100000
    sum = 0
    for _, v in totals.items():
        if v <= max_size:
            sum += v

    print(sum)

    disk_size = 70_000_000
    update_size = 30_000_000

    free_space = disk_size - totals['/']
    min_space_needed = update_size - free_space

    smallest_dir_to_delete = math.inf
    for _, v in totals.items():
        if v > min_space_needed:
            smallest_dir_to_delete = min(smallest_dir_to_delete, v)

    print(smallest_dir_to_delete)


if __name__ == "__main__":
    run()
