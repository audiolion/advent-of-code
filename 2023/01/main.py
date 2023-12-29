with open(0) as f:
    input = f.read()
    lines = input.split("\n")

strings = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9"
}

n = 0
m = 0
for line in lines:
    n1 = None
    n2 = None
    m1 = None
    m2 = None

    for i, ch in enumerate(line):
        if ch.isdigit():
            if n1 is None:
                n1 = ch
            else:
                n2 = ch
            if m1 is None:
                m1 = ch
            else:
                m2 = ch
        for s in strings.keys():
            if s in line[i:i+1+len(s)]:
                if m1 is None:
                    m1 = strings[s]
                else:
                    m2 = strings[s]

    if n2 is None:
        n2 = n1
    if m2 is None:
        m2 = m1

    n += int(n1 + n2)
    m += int(m1 + m2)

print(n)
print(m)
