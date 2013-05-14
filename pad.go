package main 
import "fmt"
import "strings"
func main() {
	var str = "1234567890"
	fmt.Println(str)

	fmt.Println("leftPad(str, \"*\", 3) ",          leftPad(str, "*", 3))
	fmt.Println("leftPad2Len(str, \"*-\", 13) ",    leftPad2Len(str, "*-", 13))
	fmt.Println("leftPad2Len(str, \"*-\", 14) ",    leftPad2Len(str, "*-", 14))
	fmt.Println("leftPad2Len(str, \"*\", 14) ",     leftPad2Len(str, "*", 14))
	fmt.Println("leftPad2Len(str, \"*-x\", 14) ",   leftPad2Len(str, "*-x", 14))
	fmt.Println("leftPad2Len(str, \"ABCDE\", 14) ", leftPad2Len(str, "ABCDE", 14))
	fmt.Println("leftPad2Len(str, \"ABCDE\", 4) ",  leftPad2Len(str, "ABCDE", 4))

	fmt.Println("rightPad(str, \"*\", 3) ",          rightPad(str, "*", 3))
	fmt.Println("rightPad(str, \"*!\", 3) ",         rightPad(str, "*!", 3))
	fmt.Println("rightPad2Len(str, \"*-\", 13) ",    rightPad2Len(str, "*-", 13))
	fmt.Println("rightPad2Len(str, \"*-\", 14) ",    rightPad2Len(str, "*-", 14))
	fmt.Println("rightPad2Len(str, \"*\", 14) ",     rightPad2Len(str, "*", 14))
	fmt.Println("rightPad2Len(str, \"*-x\", 14) ",   rightPad2Len(str, "*-x", 14))
	fmt.Println("rightPad2Len(str, \"ABCDE\", 14) ", rightPad2Len(str, "ABCDE", 14))
	fmt.Println("rightPad2Len(str, \"ABCDE\", 4) ",  rightPad2Len(str, "ABCDE", 4))
}

//TODO convert these into a 
/*
 * leftPad and rightPad just repoeat the padStr the indicated
 * number of times
 *
 */

func leftPad(s string, padStr string, pLen int) string{
	return strings.Repeat(padStr, pLen) + s
}
func rightPad(s string, padStr string, pLen int) string{
	return s + strings.Repeat(padStr, pLen);
}

/* the Pad2Len functions are generally assumed to be padded with short sequences of strings
 * in many cases with a single character sequence
 *
 * so we assume we can build the string out as if the char seq is 1 char and then
 * just substr the string if it is longer than needed
 *
 * this means we are wasting some cpu and memory work
 * but this always get us to want we want it to be
 * 
 * in short not optimized to for massive string work
 *
 * If the overallLen is shorter than the original string length 
 * the string will be shortened to this length (substr)
 * 
 */

func rightPad2Len(s string, padStr string, overallLen int) string{
	var padCountInt int
	padCountInt = 1 + ((overallLen-len(padStr))/len(padStr))
	var retStr =  s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
func leftPad2Len(s string, padStr string, overallLen int) string{
	var padCountInt int
	padCountInt = 1 + ((overallLen-len(padStr))/len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr)-overallLen):]
}

