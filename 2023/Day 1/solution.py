numbers_record = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "1": 1,
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9
}

def update_words(character, words):
    words.append('')
    return list(map(lambda word: f"{word}{character}", words))

def get_number(words):
    number = None
    for word in words:
        if word in numbers_record:
            number = numbers_record[word]
            break
    return [word, number]

def cleanup_words(words, number):
    word_index = words.index(number)
    return words[word_index + 1:]


def get_calibration_value(line):
    first_digit = 0
    second_digit = 0
    current_string_number = '';
    ## 
    ## threerjoneonepmdjgcrjlmlqvqbpzg7three
    ## t -> th -> thr -> thre -> three
    ##      h -> hr -> hre -> hree
    ##            r -> re -> ree
    ## leightwothreeninesixtsljvdl1nflg249
    ## l -> le -> lei -> leig -> leigh -> leight
    ##   -> e -> ei -> eig -> eigh -> eight
    words = []
    for character in line:
        words = update_words(character, words)
        [actual_word, actual_number] = get_number(words)
        if actual_number:
            if not first_digit:
                first_digit = actual_number
            else:
                second_digit = actual_number
            words = cleanup_words(words, actual_word)
    
    if not second_digit:
        second_digit = first_digit
    return first_digit * 10 + second_digit



def solve(file_name): 
    calibration_sum = 0
    with open(file_name, 'r') as file:
        for line in file:
            calibration_value = get_calibration_value(line)
            print(f"{line.strip()} {calibration_value}")
            calibration_sum += get_calibration_value(line)
    return calibration_sum

if __name__ == '__main__':
    calibration_sum = solve('input.txt')
    print(calibration_sum)