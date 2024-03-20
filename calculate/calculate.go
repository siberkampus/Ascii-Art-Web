package calculate

import (
	"bufio"
	"fmt"
	"os"
)

func AralikHesapla(kelime string, filename string) []string {
	var satir []string
	satirlar := make([]string, 9)

	for _, harf := range kelime {
		asciidegeri := int(harf)

		asciidegeri = (asciidegeri-32)*9 + 2
		satir = SatirOku(filename, asciidegeri, asciidegeri+8)
		for index, value := range satir {
			satirlar[index] += value
		}
	}
	return satirlar
}

func SatirOku(filename string, baslangic, bitis int) []string {
	file, _ := os.Open(filename) // Dosya acilir
	var result []string
	defer file.Close()

	scanner := bufio.NewScanner(file)
	satirsayisi := 0
	for scanner.Scan() { // dosya içeriği okunur
		satirsayisi++
		if satirsayisi >= baslangic && satirsayisi <= bitis { // belirtilen aralıktaki satırlar alınır
			result = append(result, scanner.Text()) // belirtilen aralıktaki satırlar sonuca eklenir
		}
	}
	return result
}

func Yazdir(kelimeler []string, data map[string][]interface{}, file_name string) {

	for _, kelime := range kelimeler {
		if kelime == "" {
			fmt.Println()
		} else {
			value := AralikHesapla(kelime, file_name+".txt")
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
