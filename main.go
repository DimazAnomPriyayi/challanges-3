/** * Created with VS Code. * User: Dimas Anom Priyayi * Date: 13/03/23 * Time: 19:23  * ID Reg : 1955617840-940 * Challanges 3 **/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Input text
	text := "selamat malam"

	// Hitung text dan outputkan text
	counter := make(map[string]int)
	for _, c := range text {
		fmt.Printf("%v\n", string(c))
		counter[string(c)]++
	}

	// Hitung banyak text
	chars := make([]string, 0, len(counter))
	for chr := range counter {
		chars = append(chars, chr)
	}

	// Sorting secara alfabetis
	sort.Strings(chars)

	// Tampilkan hasil perhitungan
	fmt.Printf("map [ ")
	for _, chr := range chars {
		fmt.Printf("%v: %v ", chr, counter[chr])
	}
	fmt.Printf("]")
}
