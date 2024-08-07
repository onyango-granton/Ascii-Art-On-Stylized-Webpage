## Ascii Art Output
The ASCII Art Output is a program written to generate ASCII art from a given string or its specified substring and writes the output to a file.
The format of the input received as arguments is:
```
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard
```
## NOTE
1. This program builds on the ASCII Art from  **[ASCII-REPOSITORY](https://learn.zone01kisumu.ke/git/wyonyango/ascii-art.git)**
2. This program supports ASCII Art Color from **[ASCII-ART-COLOR-REPOSITORY](https://learn.zone01kisumu.ke/git/lakoth/ascii-art-color)**
3. This program also supports ASCII-Art-fs from  **[ASCII-ART-FS REPOSITOTY](https://learn.zone01kisumu.ke/git/wyonyango/ascii-art-fs.git)**

## Usage
To use the ASCII Art Generator, follow these steps:
1. Clone this repository to your local machine using the comand below.
``` bash
git clone https://learn.zone01kisumu.ke/git/lakoth/ascii-art-output.git
```

2. Navigate to the directory where the repository is cloned.
```bash
cd ascii-art-output
```

3. Initialize your module to get the required dependancies :Run the command below to initialize your module
```bash
go mod init ascii-art-output
```
4. Run the program from the command line, providing a string as an argument and file to store art output


## Example 1

```console
go run . --output=banner.txt "hello" standard
cat -e banner.txt
```
- Output
``` 
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
```

## Example 2

```console
go run . --output=banner.txt 'Hello There!' shadow
cat -e banner.txt
```
- Output
``` 
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
```

## Dependencies
This program requires Go (Golang) to be installed on your system. You can download and install it from the [official Go website](https://golang.org/dl/).

## Contributing
Contributions to this project are welcome! If you'd like to contribute, please fork the repository and submit a pull request with your changes.

## Contributors
<body>
<div style="display: flex !important; justify-content: center !important;">
    <div style="margin: 10px;">
        <img src="images/wkycliffe.png" style="border-radius: 50% !important; width: 150px !important; height: 150px; !important" alt="Wykliffe" >
        <p style="text-align: center;">Wycliffe</p>
    </div>
    <div style="margin: 10px;">
        <img src="images/lakoth.png" style="border-radius: 50% !important; width: 150px !important; height: 150px; !important" alt="Lorna" >
        <p style="text-align: center;">Lorna</p>
    </div>
    <div style="margin: 10px;">
        <img src="images/gonyango.png" style="border-radius: 50% !important; width: 150px !important; height: 150px; !important" alt="Granton">
        <p style="text-align: center;">Granton</p>
    </div>
</div>
</body>


## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.



