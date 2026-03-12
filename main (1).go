package main

import (
	"fmt"
	"os"
	"strings"

	// "io"
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	
	
)

func main(){


	if len(os.Args) != 3 {  // check if the Argument inputed in the terminal is exactly three
		fmt.Println("Error!: user are required to provide 3 argument for the program")

	}

	rawFile := os.Args[1] // assign the name of the second argument to this variable
	newFile := os.Args[2] // assign the name of the third argument to this variable

	file1, err1 := os.ReadFile(rawFile) // read the content of the file

	if err1 != nil { // check if there was an error while reading the content of rawFile
		fmt.Println("Error", err1)
	} 

	
	word1 := strings.Fields(string(file1)) // split the content in file1 into an array of strings using whitespace as a delimiter
	
	new := tags(word1) 
	
	new = quote(new)
	new = punc(new)
	new = article(new)
	



	file := os.WriteFile(newFile,[]byte(new),0644) // write the content of new into the file newFile 
	if file != nil { // check if there was an error while writing
		fmt.Println("Error",file)

	}

	 


}

func tags(a []string )(string) {
	var b []string


	for i:= 0; i < len(a); i++ {
		if i != 0 && (a[i] != "(hex)" || a[i] != "(bin)" || a[i] != "(up)" || a[i] != "(cap)" || a[i] != "(low)" || a[i] != "(up," || a[i] != "(cap," || a[i] != "(low,")  { // check if the tags are not the first index
			switch a[i] {
				case "(hex)":
					val, _ := strconv.ParseInt(a[i-1],16,64) // convert string into number and convert from the declare base to decimal
					b[len(b)-1] = strconv.FormatInt(val, 10) // convert integer to string
				case  "(bin)" :
					val,_ := strconv.ParseInt(a[i-1], 2, 64)
					b[len(b)-1] = strconv.FormatInt(val, 10)
				case "(up)":
					val := strings.ToUpper(a[i-1]) // convert string to uppercase 
					b[len(b)-1] = val
				case "(low)" :
					val := strings.ToLower(a[i-1]) // convert string to lowercase
					b[len(b)-1] = val
				case "(up,":
					val,_ := strconv.ParseInt(strings.Trim(a[i+1],")"),10,64) //trim ) out 
					tex := strings.ToUpper(a[i-1])
					b[len(b)-1] = tex
					for j:= 2; val >= int64(j) ; j++ {
						b[len(b)-int(j)] = strings.ToUpper(a[i-int(j)]) 
					}
					i++
				case "(low,":
					val,_ := strconv.ParseInt(strings.Trim(a[i+1],")"),10,64)
					tex := strings.ToLower(a[i-1])
					b[len(b)-1] = tex
					for j:= 2; val >= int64(j) ; j++ {
						b[len(b)-int(j)] = strings.ToLower(a[i-int(j)])
					}
					i++
				case "(cap)" :
					caser := cases.Title(language.English) // declare the language 
					val := caser.String(a[i-1]) // capitalize the string
 					b[len(b)-1] = val
				case "(cap," :
					caser := cases.Title(language.English)
		 			val := caser.String(a[i-1]) 
		 			b[len(b)-1] = val
		 			num,_ := strconv.ParseInt(strings.Trim(a[i+1],")"),10,64)
		 			for j:= 2; num >= int64(j) ; j++ {
		 				b[len(b)-int(j)] = caser.String(a[i-int(j)])
		 			}
		 			i++

				default:
					b = append(b,a[i])
			}

		} else {
			b = append(b,a[i])
		}

		
	}
	newWord := strings.Join(b, " ")
	return newWord



}

func punc(s string) string {
	punctuations := []string{".", ",", "!", "?", ":", ";"} 

	for _, p := range punctuations {
		for strings.Contains(s, " "+p) {
			s = strings.Replace(s, " "+p, p, 1)
		}
	}

	result := []rune(s)
	var out []rune
	puncSet := map[rune]bool{'.': true, ',': true, '!': true, '?': true, ':': true, ';': true}

	for i := 0; i < len(result); i++ {
		out = append(out, result[i])
		if puncSet[result[i]] {
			if i+1 < len(result) && !puncSet[result[i+1]] && result[i+1] != ' ' && result[i+1] != '\'' {
				out = append(out, ' ')
			}
		}
	}
	return string(out)
}

func quote(s string) string {
	words := strings.Fields(s)
	var result []string
	inQuote := false

	for i := 0; i < len(words); i++ {
		if words[i] == "'" && !inQuote {
			inQuote = true
			if i+1 < len(words) {
				words[i+1] = "'" + words[i+1]
			} else {
				result = append(result, "'")
			}
		} else if words[i] == "'" && inQuote {
			inQuote = false
			if len(result) > 0 {
				result[len(result)-1] = result[len(result)-1] + "'"
			} else {
				result = append(result, "'")
			}
		} else {
			result = append(result, words[i])
		}
	}
	return strings.Join(result, " ")
}

func article(s string) string {
	words := strings.Fields(s)
	vowels := "aeiouhAEIOUH"

	for i := 0; i < len(words)-1; i++ {
		lower := strings.ToLower(words[i])
		if lower == "a" {
			nextWord := words[i+1]
			checkWord := strings.TrimLeft(nextWord, "'\"")
			if len(checkWord) > 0 && strings.ContainsRune(vowels, rune(checkWord[0])) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}
	return strings.Join(words, " ")
}