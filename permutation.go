package main

import (
	"strings"
	"golang.org/x/net/idna"
	"github.com/likexian/whois"
)

var p *idna.Profile
	
func permute(domain string){
	dico := dicoSetup()
	res := []string{}
	var asciiResult string
	word := strings.Split(domain, ".")
	splittedWorld := strings.Split(word[0], "")
	p = idna.New()
	for i, v := range(splittedWorld) {
		for _, altChar := range(dico[v]){
			asciiResult, _ = p.ToASCII(strings.Join(splittedWorld[:i], "")+altChar+strings.Join(splittedWorld[i+1:], "")+"."+word[1])
			if stringInSlice(asciiResult, res){
				debugmsg("Duplicate string: " + asciiResult)
			}else{
				res = append(res, asciiResult)
			}
			
		}
	}
	returnResult(res)
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func returnResult(resultList []string) {
	var textToWrite, registrar string
	var isRegister bool
	for _, result := range(resultList) {
		isRegister, registrar = asWhois(result)
		if isRegister {
			textToWrite = result + ",True," + registrar 
		}else{
			textToWrite = result + ",False," 
		}
		writeInFile(textToWrite, outputfile)
	}
}

func asWhois(d string) (bool, string) {
	result, err := whois.Whois(d)
	if !warningerr(err) {
		substr := "No match for"
		debugmsg(result)
		if strings.Contains(result, substr){
			return false, ""
		}else{
			return true, getRegistrar(result)
		}
	}
	errormsg("Whois error for \"" + d + "\"")
	return false, ""
}

func getRegistrar(whoisResult string) string {
	regisSubstr := "Registrar:"
	splittedResult := strings.Split(whoisResult, "\n")
	for _,line := range(splittedResult){
		if strings.Contains(line, regisSubstr){
			return strings.Split(line,":")[1]
		}
	}
	return ""
}

func dicoSetup() map[string][]string {
	m := make(map[string][]string)
	m["a"] = []string{"a", "ɑ", "4", "α", "∂", "а", "А", "Д", "д", "Ѧ", "ѧ", "Å", "Ä", "Ã", "Â", "Á", "À", "à", "á", "â", "ã", "ä", "å", "Δ", "Α", "Λ", "λ", "∧", "∂", "a̰", "ă", "ā"}
	m["b"] = []string{"b", "8", "6", "13", "ß", "б", "Б", "в", "В", "Ь", "ь", "Β", "Ѣ", "ѣ", "b̤", "b̰"}
	m["c"] = []string{"c", "¢", "©", "С", "с", "Ç", "ç", "⊂", "⊆", "⟨", "<", "‹", "č", "ć"}
	m["d"] = []string{"d", "0", "Ð", "đ", "d̼", "d̻", "d̺", "d̪"}
	m["e"] = []string{"e", "€", "£", "ë", "Є", "є", "Е", "е", "З", "з", "È", "É", "Ê", "Ë", "è", "é", "ê", "Ε", "Σ", "ξ", "ε", "Ё", "ё", "ẽ", "e̽", "ë", "e̙", "e̘", "e̝", "e̞", "ē"}
	m["f"] = []string{"f", "ph", "ƒ", "∫", "ʄ"}
	m["g"] = []string{"g", "6", "&", "C-", "ϑ", "ɕ", "ɠ", "ġ"}
	m["h"] = []string{"h", "#", "Η", "И", "и"}
	m["i"] = []string{"i", "ì", "Ì", "1", "¡", "І", "і", "Ї", "ї", "Î", "Í", "í", "î", "Ι", "Ї", "ї", "ɨ", "i̠", "l", "ī"}
	m["j"] = []string{"j", ";", "ʝ", "Ј", "ј"}
	m["k"] = []string{"k", "X", "ɮ", "К", "к", "Κ", "κ"}
	m["l"] = []string{"l", "1", "I", "£", "1_", "ℓ", "ι", "⌊", "ɭ", "ǀ"}
	m["m"] = []string{"m", "М", "м", "Μ", "ɱ", "ʍ"}
	m["n"] = []string{"n", "₪", "Л", "л", "П", "п", "Ñ", "Ν", "η", "π", "ℵ", "∩", "И", "и", "ग"}
	m["o"] = []string{"o", "0", "¤", "°", "О", "о", "Ѳ", "ѳ", "Ф", "Ò", "Ó", "Ô", "Õ", "Ö", "ø", "ö", "õ", "ô", "ó", "ò", "ð", "Θ", "Ο", "Ω", "θ", "σ", "∅", "◊", "⊗", "⊕", "δ", "º", "ʘ",  "ō", "ŏ"}
	m["p"] = []string{"p", "9", "¶", "Р", "р", "Ρ", "ρ"}
	m["q"] = []string{"q"}
	m["r"] = []string{"r", "®", "ʁ", "Я", "я", "Г", "г", "ℜ", "r̥", "r̪", "ṛ"}
	m["s"] = []string{"s", "5", "$", "z", "§", "Ѕ", "ѕ", "ς", "Š", "š", "ऽ"}
	m["t"] = []string{"t", "7", "+", "1", "†", "Т", "т", "Γ", "τ", "†"}
	m["u"] = []string{"u", "v", "µ", "J", "Ц", "ц", "Ù", "Ú", "Û", "Ü", "υ", "∪", "ü", "û", "ú", "ù", "Џ", "џ", "ʊ", "u̟", "ū"} 
	m["v"] = []string{"v", "Ѵ", "ѵ", "∇", "√", "∨", "∀"}
	m["w"] = []string{"w", "vv", "VV", "Щ", "щ", "Ш", "ш", "Ѡ", "ѡ", "ϖ", "ω", "ɯ", "w"}
	m["x"] = []string{"x", "×", "Х", "х", "Ж", "ж", "χ"}
	m["y"] = []string{"y", "j", "Ѱ", "ѱ", "Ψ", "φ", "¥", "Ч", "ч", "Ý", "ÿ", "γ", "ϒ", "Ÿ", "Џ", "џ"}
	m["z"] = []string{"z","2", "7_", "Ζ", "ζ", "Ꙁ", "ꙁ", "ž"}
	m["."] = []string{".","_", "-", ""}
	m["-"] = []string{"-", "_", ".", ""}
	m["_"] = []string{"_", "-", ".", ""}
	m["0"] = []string{"0", ""}
	m["1"] = []string{"1", ""}
	m["2"] = []string{"2", ""}
	m["3"] = []string{"3", ""}
	m["4"] = []string{"4", ""}
	m["5"] = []string{"5", ""}
	m["6"] = []string{"6", ""}
	m["7"] = []string{"7", ""}
	m["8"] = []string{"8", ""}
	m["9"] = []string{"9", ""}
	
	return m
}