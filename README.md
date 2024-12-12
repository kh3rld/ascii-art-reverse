# ASCII Art Reverse

The ASCII Art Reverse is designed to reverse the process of converting a art into its normal string format.

## Features

- Accepts a command-line flag `--reverse=<fileName>` to specify the input file containing ASCII art.
- Prints the decoded string to the console.
- Supports additional options and banners if other ASCII art projects are implemented.
- Can also run with a single string argument.

## Usage

To use the program, you must provide the `--reverse` flag followed by the name of the file containing the ASCII art representation. The program will then output the corresponding string.

### Command Format

```bash
go run . --reverse=<fileName>
```

### Example

1. Create a text file named `file.txt` with the following content:

```
 _              _   _
| |            | | | |
| |__     ___  | | | |   ___
|  _ \   / _ \ | | | |  / _ \
| | | | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/


```

2. Run the program with the following command:

```bash
go run . --reverse=file.txt
```

3. The output will be:

```
hello
```

## Additional Information

- If the flag is not formatted correctly, the program will return the following usage message:

```
Usage: go run . [OPTION]

EX: go run . --reverse=<fileName>
```

## Installation

Ensure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

Clone the repository and navigate to the project directory:

```bash
git clone https://learn.zone01kisumu.ke/git/doonyango/ascii-art-reverse 
cd ascii-art-reverse
```

Then, you can run the program using the command provided in the usage section.

## Contributing

Contributions are welcome! If you have suggestions for improvements or additional features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
