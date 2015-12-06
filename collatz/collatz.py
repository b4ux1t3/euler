def main():
    highest = 0
    highestInput = 0

    result = 0

    for i in range(1000000):
        result = collatz(i + 1)

        if result > highest:
            highest = result
            highestInput = i + 1

    print "Most steps:", highestInput, highest

def collatz(input):
    count = 0

    while input > 1:
        if input % 2 == 0:
            input /= 2
            count += 1
        else:
            input = (input * 3) + 1
            count += 1

    count += 1
    return count


if __name__ == "__main__":
    main()

