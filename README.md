# Ascii-art
ASCII-art is a program which consists in receiving a string as an argument and outputting the string in a graphical representation using ASCII.

## To run the program
Simply copy the below command and change the text between the "" symbol. And make sure that you are in the **App** directory.

- $ go run . "enter-yours-text" BannerName

    ex: go run . "Hello There!" shadow | cat -e

Or try the test file.
- $ ../Testing/Test.sh
- $ ../Testing/BannaerTest.sh

Also you can save the art in a txt file by using thr follwing command (Make sure to use the same format as following):
- $ go run . --output="FileName.txt" "enter-yours-text" BannerName

## Currently Supported Banner names
- [x] standard
- [x] Thinkertoy banner
- [x] shadow banner
- [x] output banner

## Future updates 
- [ ] reverse text
- [ ] colored art
- [ ] and many more...

