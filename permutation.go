package main

import (
	"strings"
	"golang.org/x/net/idna"
	"github.com/likexian/whois"
	"strconv"
)

var p *idna.Profile
var finalResult []string
var maxRecursiveLevel int
var cpt = 0
var cpt_b = 0
	
func permute_string(domain string){
	maxRecursiveSetup(domain)
	finalResult := generator(domain)
	debugmsg("permute_string - first list generated")
	permute_list(finalResult)
	debugmsg("permute_string - permutation ended ... file will be writen soon")
	returnResult(listToPunycode(finalResult))
}

func permute_list(domains []string){
	debugmsg("permute_list - function called")
	debugmsg("permute_list - cpt bis=" + strconv.Itoa(cpt))
	cpt = cpt + 1
	if cpt <= maxRecursiveLevel {
		debugmsg("permute_list - enter IF")
		var workList []string
		debugmsg("permute_list - input length=" + strconv.Itoa(len(domains)))
		for _, val := range(domains) {
			debugmsg("permute_list - enter FOR")
			workList = generator(val)
			debugmsg("permute_list - list generated")
			finalResult = mergeList(finalResult, workList)
			debugmsg("permute_list - list merged")
			permute_list_bis(workList)
			cpt_b = 0
		}
	}else{
		debugmsg("permute_list - enter ELSE")
		debugmsg("permute_list - recursion ended")
	}
}

func permute_list_bis(domains []string){
	debugmsg("permute_list_bis - function called")
	debugmsg("permute_list_bis - cpt bis=" + strconv.Itoa(cpt_b))
	cpt_b = cpt_b + 1
	if cpt_b <= maxRecursiveLevel {
		debugmsg("permute_list_bis - enter IF")
		var workList []string
		debugmsg("permute_list_bis - input length=" + strconv.Itoa(len(domains)))
		for _, val := range(domains) {
			debugmsg("permute_list_bis - enter FOR")
			workList = generator(val)
			debugmsg("permute_list_bis - list generated")
			finalResult = mergeList(finalResult, workList)
			debugmsg("permute_list_bis - list merged")
			permute_list_bis(workList)
		}
	}else{
		debugmsg("permute_list_bis - enter ELSE")
		debugmsg("permute_list_bis - recursion ended")
	}
}

func mergeList (outputList []string, toMerge []string) []string {
	debugmsg("mergeList - function called")
	for _, v := range(toMerge){
		if stringInSlice(v, outputList) {
			debugmsg("Duplicate string: " + toPunycode(v))
		}else{
			outputList = append(outputList, v)
		}
	}
	return outputList
}

func maxRecursiveSetup(s string) {
	word := strings.Split(s, ".")
	maxRecursiveLevel = len(word[0])
	debugmsg("maxRecursiveSetup - maxRecursiveLevel=" + strconv.Itoa(len(word[0])))
}

func listToPunycode(input []string) []string {
	p = idna.New()
	var output []string
	var asciiStr string
	for _, v := range(input){
		asciiStr, _ = p.ToASCII(v)
		output = append(output, asciiStr)
	}
	return output
}

func toPunycode(s string) string {
	p = idna.New()
	asciiStr, _ := p.ToASCII(s)
	return asciiStr
}

func generator(s string) []string {
	debugmsg("generator - function called")
	dico := dicoSetup()
	sl := []string{}
	var res string
	word := strings.Split(s, ".")
	splittedWorld := strings.Split(word[0], "")
	for i, v := range(splittedWorld) {
		for _, altChar := range(dico[v]) {
			res = strings.Join(splittedWorld[:i], "")+altChar+strings.Join(splittedWorld[i+1:], "")+"."+word[1]
			if stringInSlice(res, sl) {
				debugmsg("Duplicate string: " + toPunycode(res))
			}else{
				sl = append(sl, res)
			}
		}
	}
	return sl
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
	for _,line := range(splittedResult) {
		if strings.Contains(line, regisSubstr) {
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