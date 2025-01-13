import sys
import os


if __name__ == "__main__":
    top_folder = sys.argv[1]
    subfolder = sys.argv[2]
    subfolder_path = os.path.join(top_folder, subfolder)
    # create subfolder
    os.mkdir(path=subfolder_path)
    # create files in subfolder using subfolder name
    # it expects subfolder to follow the convention
    # dayX name which then makes the files inside
    # dayX.go and dayX.input.txt
    go_file_path = os.path.join(subfolder_path, f'{subfolder}.go')

    
    top_folder = top_folder.replace(".", "").replace("\\", "")
    code = f"""package main

import "github.com/joaooliveirapro/adventofcodego/utils"

func part1(input *[]string) {{

}}

func part2(input *[]string) {{

}}

func main() {{
	input := utils.ReadInput("./{top_folder}/{subfolder}/{subfolder}.input.txt")
	part1(&input)
	part2(&input)
}}"""

    with open(go_file_path, mode='w') as gofile:
        gofile.write(code)
        
    input_file_path = os.path.join(subfolder_path, f'{subfolder}.input.txt')
    with open(input_file_path, mode='w') as inputfile:
        inputfile.write("# add your input here #")
    