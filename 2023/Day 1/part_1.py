def get_calibration_value(line):
    line_length = len(line)
    string_number = ''
    print(line)
    for character in line:
        if character.isdigit():
            string_number = f"{string_number}{character}"
    
    print(string_number)
    first_digit = string_number[0]
    second_digit = string_number[-1]
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