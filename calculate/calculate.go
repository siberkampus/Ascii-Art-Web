package calculate

import (
	"bufio"
	"fmt"
	"os"
)

func CalcSpace(word string, filename string) []string {
	var line []string
	lines := make([]string, 9)

	for _, letter := range word {
		asciivalue := int(letter)

		asciivalue = (asciivalue-32)*9 + 2
		line = ReadLine(filename, asciivalue, asciivalue+8)
		for index, value := range line {
			lines[index] += value
		}
	}
	return lines
}

func ReadLine(filename string, start, end int) []string {
	file, _ := os.Open(filename) // Dosya acilir
	var result []string
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linenum := 0
	for scanner.Scan() { // dosya içeriği okunur
		linenum++
		if linenum >= start && linenum <= end { // belirtilen aralıktaki satırlar alınır
			result = append(result, scanner.Text()) // belirtilen aralıktaki satırlar sonuca eklenir
		}
	}
	return result
}

func Print(words []string, data map[string][]interface{}, file_name string) {

	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			value := CalcSpace(word, file_name+".txt")
			for j, i := range value {
				if len(value)-1 != j {
					if data["result"] == nil {
						data["result"] = make([]interface{}, 0)
					}
					data["result"] = append(data["result"], fmt.Sprintf("%s\n", i))
				}
			}
		}
	}

}
