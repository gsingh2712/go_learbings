package main

import (
	"fmt"
	"unicode/utf8"
)

/*
This String contains Special Characters / utf8 encoded chars,
these chars take more than 1 byte
A simple for loop reading text[i] will not work , since it will read incremently just 1 byte from the index

Hence use utf8.DecodeRuneInString(), because it gives the size of char rune in question.
and use it to read next character.

Side Note :-
1. You can count the runes/utf8 chars in text via utf8.RuneCountInString(str)

2. A rune is typeless values and its default type is rune and can be converted to any integer types
and depending on context it can take 1byte , 2byte , 3byte , 4byte
Hence in case of string with utf8 chars rune is prerefered to read each char.

*/
func main() {
	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.

Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.

Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	r, size := utf8.DecodeRuneInString("öykü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ykü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("kü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	r, size = utf8.DecodeRuneInString("ü")
	fmt.Printf("rune: %c size: %d bytes.\n", r, size)

	// for range loop automatically decodes the runes
	//   but it gives you the position of the current rune
	//   instead of its size.

	// for _, r := range text {}
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)

		i += size
	}
	fmt.Println()
}
