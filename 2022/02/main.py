decoder = {
    "A": "rock",
    "B": "paper",
    "C": "scissors",

    "X": "rock",
    "Y": "paper",
    "Z": "scissors",
}

shape_score = {
    "X": 1,
    "Y": 2,
    "Z": 3,
}

outcome_score = {
    "X": 0,
    "Y": 3,
    "Z": 6,
}

combos = {
    "A X": 3,
    "A Y": 6,
    "A Z": 0,
    "B X": 0,
    "B Y": 3,
    "B Z": 6,
    "C X": 6,
    "C Y": 0,
    "C Z": 3,
}

combos_2 = {
    "A X": 3,
    "A Y": 1,
    "A Z": 2,
    "B X": 1,
    "B Y": 2,
    "B Z": 3,
    "C X": 2,
    "C Y": 3,
    "C Z": 1,
}


def run():
    with open("in.1", "r") as f:
        input = f.read()

    lines = input.split("\n")

    score = 0
    for line in lines:
       score += combos[line] + shape_score[line.split(" ")[1]]

    print(score)

    score_2 = 0
    for line in lines:
        score_2 += outcome_score[line.split(" ")[1]] + combos_2[line]

    print(score_2)

if __name__ == "__main__":
    run()
