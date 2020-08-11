package transformations
import (
	"strings"
	
)

/*
https://github.com/SpiderLabs/ModSecurity/blob/b66224853b4e9d30e0a44d16b29d5ed3842a6b11/src/actions/transformations/cmd_line.cc
Copied from modsecurity
deleting all backslashes [\]
deleting all double quotes ["]
deleting all single quotes [']
deleting all carets [^]
deleting spaces before a slash /
deleting spaces before an open parentesis [(]
replacing all commas [,] and semicolon [;] into a space
replacing all multiple spaces (including tab, newline, etc.) into one space
transform all characters to lowercase
*/
func CmdLine(data string) string{
	space := 0
	ret := ""
	for _, a:= range data{
        switch (a) {
            /* remove some characters */
            case '"':
            case '\'':
            case '\\':
            case '^':
                break;

            /* replace some characters to space (only one) */
            case ' ':
            case ',':
            case ';':
            case '\t':
            case '\r':
            case '\n':
                if (space == 0) {

                    ret += " "
                    space++
                }
                break

            /* remove space before / or ( */
            case '/':
            case '(':
                if (space != 0) {
                	ret = ret[0:len(ret)-2] //TODO: CHECK
                    //ret.pop_back();
                }
                space = 0;

                ret += string(a)
                break;

            /* copy normal characters */
            default :
                b := strings.ToLower(string(a))
                ret += b
                space = 0
                break
        }
    }
	return data
}