from copy import copy

decoder = {
    "A": "rock",
    "B": "paper",
    "C": "scissors",

    "X": "rock",
    "Y": "paper",
    "Z": "scissors",
}


def run():
    with open("in.1", "r") as f:
        input = f.read()

    lines = input.split("\n")

    def total(a, b):
        points = b.points
        if a < b:
            points += 6
        elif a == b:
            points += 3
        return points

    score = 0
    for line in lines:
        a, b = [Shape(decoder[s]) for s in line.split(" ")]
        score += total(a, b)

    print(score)

    score_2 = 0
    for line in lines:
        parts = line.split(" ")
        a = Shape(decoder[parts[0]])
        s = Strategy(parts[1])
        b = s.shape_choice(a)
        score_2 += total(a, b)

    print(score_2)


class Strategy:
    outcome = None

    def __init__(self, enc):
        if enc == "X":
            self.outcome = "lose"
        elif enc == "Y":
            self.outcome = "draw"
        elif enc == "Z":
            self.outcome = "win"

    def __str__(self):
        return self.outcome.title()

    def shape_choice(self, a):
        if self.outcome == "draw":
            return copy(a)
        if self.outcome == "lose":
            return Shape(a.strength())
        if self.outcome == "win":
            return Shape(a.weakness())


class Shape:
    name = None
    points = None

    def __init__(self, name):
        self.name = name
        if self.name == "rock":
            self.points = 1
        elif self.name == "paper":
            self.points = 2
        elif self.name == "scissors":
            self.points = 3

    def __eq__(self, other):
        return self.name == other.name

    def __lt__(self, other):
        if self.name == "rock":
            return other.name == "paper"
        if self.name == "paper":
            return other.name == "scissors"
        if self.name == "scissors":
            return other.name == "rock"

    def __gt__(self, other):
        if self.name == "rock":
            return other.name == "scissors"
        if self.name == "paper":
            return other.name == "rock"
        if self.name == "scissors":
            return other.name == "paper"

    def __str__(self):
        return self.name.title()

    def weakness(self):
        if self.name == "rock":
            return "paper"
        if self.name == "paper":
            return "scissors"
        if self.name == "scissors":
            return "rock"

    def strength(self):
        if self.name == "rock":
            return "scissors"
        if self.name == "paper":
            return "rock"
        if self.name == "scissors":
            return "paper"


if __name__ == "__main__":
    run()
