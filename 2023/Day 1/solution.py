
def get_calibration_value(line):
    line_length = len(line)
    index = 0;
    first_digit = ''
    second_digit = ''
    while index < line_length/2:
        start_character = line[index]
        end_character = line[line_length - (index + 1)]
        if not first_digit and start_character.isdigit():
            first_digit = start_character
        if not second_digit and end_character.isdigit():
            second_digit = end_character
        index += 1
    if not first_digit:
        first_digit = second_digit    
    if not second_digit:
        second_digit = first_digit        
    return int(f"{first_digit}{second_digit}")



def solve(file_name): 
    calibration_sum = 0
    with open(file_name, 'r') as file:
        for line in file:
            calibration_value = get_calibration_value(line)
            print(f"{line} {calibration_value}")
            calibration_sum += get_calibration_value(line)
    return calibration_sum

if __name__ == '__main__':
    calibration_sum = solve('input.txt')
    print(calibration_sum)