new_list_line_file = []
with open('example.txt', 'r') as file:
    lines = file.readlines()
    for i, line in enumerate(lines):
         new_list_line_file.append(str(i) + ": " + lines[i])

with open('example.txt_new', 'w') as file:
    file.writelines(new_list_line_file)
        