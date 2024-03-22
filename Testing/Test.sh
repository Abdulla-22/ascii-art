echo Test 1:
echo "$ go run . \"\" | cat -e"
go run . "" | cat -e

echo -e "\n"

echo Test 2:
echo "$ go run . \"\n\" | cat -e"
go run . "\n" | cat -e

echo -e "\n"

echo Test 3:
echo "$ go run . \"Hello\n\" | cat -e"
go run . "Hello\n" | cat -e



echo Test 4:
echo "$ go run . \"hello\" | cat -e"
go run . "hello" | cat -e



echo Test 5:
echo "$ go run . \"HeLlO\" | cat -e"
go run . "HeLlO" | cat -e


echo Test 6:
echo "$ go run . \"Hello There\" | cat -e"
go run . "Hello There" | cat -e


echo Test 7:
echo "$ go run . \"1Hello 2There\" | cat -e"
go run . "1Hello 2There" | cat -e


echo Test 8:
echo "$ go run . \"{Hello There}\" | cat -e"
go run . "{Hello There}" | cat -e


echo Test 9:
echo "$ go run . \"Hello\nThere\" | cat -e"
go run . "Hello\nThere" | cat -e


echo Test 10:
echo "$ go run . \"Hello\n\nThere\" | cat -e"
go run . "Hello\n\nThere" | cat -e


echo Test 11:
echo "$ go run . \"HELLO\" | cat -e"
go run . "HELLO" | cat -e


echo Test 12:
echo "$ go run . \"HeLlo HuMaN\" | cat -e"
go run . "HeLlo HuMaN" | cat -e


echo Test 13:
echo "$ go run . \"{Hello \& There \#}\" | cat -e"
go run . "{Hello & There #}" | cat -e


echo Test 14:
echo "$ go run . 'hello There 1 to 2!' | cat -e"
go run . 'hello There 1 to 2!' | cat -e


echo Test 15:
echo "$ go run . \"MaD3IrA\&LiSboN\" | cat -e"
go run . "MaD3IrA&LiSboN" | cat -e


echo Test 16:
echo "$ go run . \"1a\"#FdwHywR\&\/\(\)=\" | cat -e"
go run . "1a\"#FdwHywR&/()=" | cat -e


echo Test 17:
echo "$ go run . \"{\|}\~\" | cat -e"
go run . "{|}~" | cat -e


echo -e "\n"

echo Test 18:
echo "$ go run . \"[\\]^_ \'a\" | cat -e"
go run . "[\]^_ 'a" | cat -e



echo Test 19:
echo "$ go run . \"RGB\" | cat -e"
go run . "RGB" | cat -e


echo Test 20:
echo "$ go run . \"\:\;\<\=\>\?\@\" | cat -e"
go run . ":;<=>?@" | cat -e


echo Test 21:
echo "$ go run . '\!" \#\$%\&'"'"'()*+,-./' | cat -e"
go run . '\!" #$%&'"'"'()*+,-./' | cat -e


echo Test 22:
echo "$ go run . \"ABCDEFGHIJKLMNOPQRSTUVWXYZ\" | cat -e"
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e


echo Test 23:
echo "$ go run . \"abcdefghijklmnopqrstuvwxyz\" | cat -e"
go run . "abcdefghijklmnopqrstuvwxyz" | cat -e