// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"ascii-art-color/printingasciipackage"
// 	"ascii-art-color/utils"
// )

// func main() {
// 	utils.FileValidation("standard.txt", 6623)
// 	utils.FileValidation("thinkertoy.txt", 5558)
// 	utils.FileValidation("shadow.txt", 7463)
// 	utils.FileValidation("groupBanner.txt", 10871)
// 	switch banner {
// 	case "standard":
// 		ap := printingasciipackage.PrintingAscii(os.Args[2], "standard.txt", "\033[0m", "")
// 		if strings.HasSuffix(os.Args[1], ".txt") && !utils.IsValid(strings.TrimPrefix(os.Args[1], "--output=")) && utils.IsValidExtension(strings.TrimPrefix(os.Args[1], "--output=")) {
// 			err := os.WriteFile(strings.TrimPrefix(os.Args[1], "--output="), []byte(ap), 0o644)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		} else {
// 			fmt.Fprintln(os.Stderr, `Error: output file should be [filename].txt and have a non-banner name`)
// 		}

// 	case "thinkertoy":
// 		ap := printingasciipackage.PrintingAscii(os.Args[2], "thinkertoy.txt", "\033[0m", "")
// 		if strings.HasSuffix(os.Args[1], ".txt") && !utils.IsValid(strings.TrimPrefix(os.Args[1], "--output=")) && utils.IsValidExtension(strings.TrimPrefix(os.Args[1], "--output=")) {
// 			err := os.WriteFile(strings.TrimPrefix(os.Args[1], "--output="), []byte(ap), 0o644)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		} else {
// 			fmt.Fprintln(os.Stderr, `Error: output file should be [filename].txt and have a non-banner name`)
// 		}
// 	case "shadow":
// 		ap := printingasciipackage.PrintingAscii(os.Args[2], "shadow.txt", "\033[0m", "")
// 		if strings.HasSuffix(os.Args[1], ".txt") && !utils.IsValid(strings.TrimPrefix(os.Args[1], "--output=")) && utils.IsValidExtension(strings.TrimPrefix(os.Args[1], "--output=")) {
// 			err := os.WriteFile(strings.TrimPrefix(os.Args[1], "--output="), []byte(ap), 0o644)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		} else {
// 			fmt.Fprintln(os.Stderr, `Error: output file should be [filename].txt and have a non-banner name`)
// 		}
// 	case "groupbanner":
// 		ap := printingasciipackage.PrintingAscii(os.Args[2], "groupBanner.txt", "\033[0m", "")
// 		if strings.HasSuffix(os.Args[1], ".txt") && !utils.IsValid(strings.TrimPrefix(os.Args[1], "--output=")) && utils.IsValidExtension(strings.TrimPrefix(os.Args[1], "--output=")) {
// 			err := os.WriteFile(strings.TrimPrefix(os.Args[1], "--output="), []byte(ap), 0o644)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		} else {
// 			fmt.Fprintln(os.Stderr, `Error: output file should be [filename].txt and have a non-banner name`)
// 		}
// 	default:
// 		fmt.Fprintln(os.Stderr, `Usage: go run . [OPTION] [STRING] [BANNER]

// EX: go run . --output=<fileName.txt> something standard`)
// 		return
// 	}
// 	}
package main