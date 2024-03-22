# ASCII Art

This project is a program written in Go that takes a string as an argument and outputs the string in a graphical representation using ASCII characters. The goal is to create a visual representation of the input string using ASCII art.

## Usage

To run the program, use the following command:

Simply copy the below command and change the text between the "" symbol. And make sure that you are in the **App** directory.

```shell
go run . "enter-yours-text" "BannerName"
```
```shell
    ex: go run . "Hello There!" "shadow" | cat -e
```
Or try the test file.

```shell
 ../Testing/Test.sh
 ../Testing/BannaerTest.sh
```
Also you can save the art in a txt file by using thr follwing command (Make sure to use the same format as following):
```shell
- $ go run . --output="FileName.txt" "enter-yours-text" BannerName
```


<!-- ```shell
go run . <input_string>
```

Replace `<input_string>` with the string you want to convert to ASCII art. -->

## Banner Files

The project includes several banner files with predefined graphical templates represented using ASCII characters. These files are formatted in a way that does not require any changes.

- `shadow`
- `standard`
- `thinkertoy`

## Banner Format

Each character in the banner files has a height of 8 lines. Characters are separated by a new line (`\n`). Here is an example of the characters ' ', '!', and '"':

```
......
......
......
......
......
......
......
......

._..
|.|.
|.|.
|.|.
|_|.
(_).
....
....

._._..
(.|.).
.V.V..
......
......
......
......
......
```


**Note:** Please refer to the ASCII manual for more information on ASCII characters and their representations.

## Currently Supported Banner names

- [x] standard
- [x] Thinkertoy banner
- [x] shadow banner
- [x] output banner

## Future updates 

- [ ] reverse text
- [ ] colored art
- [ ] and many more...

