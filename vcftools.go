package vcftools

import "strings"

func VariantInLine(infoLs []string) bool {
    for _, value := range infoLs {
        gt := strings.Split(value, ":")[0]
        if gt != "./." && gt != "0/0" {
	    return true
	}
    }
    return false
}