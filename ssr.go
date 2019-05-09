package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func main() {
	http.HandleFunc("/", ssr)
	http.HandleFunc("/dxxzst", dxxzst)

	http.ListenAndServe(":8922", nil)

	//str := httpGet("https://github.com/dxxzst/Free-SS-SSR")
	//reg := httpRegex(str)
	//result := strings.Join(reg, "\n")
	//strbytes := []byte(result)
	//fmt.Println(base64.StdEncoding.EncodeToString(strbytes))
}

func ssr(w http.ResponseWriter, r *http.Request) {
	ssrshare,_ := ssrshare("https://raw.githubusercontent.com/ImLaoD/sub/master/ssrshare.com")
	fmt.Fprintf(w, strings.Join(ssrshare, "\n"))
}

func dxxzst(w http.ResponseWriter, r *http.Request) {
	dxxzst_url := "https://github.com/dxxzst/Free-SS-SSR"
	ssrshare_url := "https://raw.githubusercontent.com/ImLaoD/sub/master/ssrshare.com"
	liesauer_url := "https://www.liesauer.net/yogurt/subscribe?ACCESS_TOKEN=DAYxR3mMaZAsaqUb"

	str, err := httpGet(dxxzst_url)
	if err != nil {
		fmt.Sprint(w, err)
	}
	reg, err := httpRegex(str)

	ssrshareStr,_ := ssrshare(ssrshare_url)
	liesauerStr,_ := ssrshare(liesauer_url)
	reg = append(append(reg, ssrshareStr...), liesauerStr...)
	sort.Strings(reg)
	result := strings.Join(RemoveDuplicatesAndEmpty(reg), "\n")
	strbytes := []byte(result)
	fmt.Fprintf(w, base64.StdEncoding.EncodeToString(strbytes))
}
func ssrshare(url string) (resultstr []string,err error){
	str, err := httpGet(url)
	if err != nil {
		return resultstr, err
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	return strings.Split(string(decodeBytes),"\n"),err
}

func httpGet(url string) (str string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		fmt.Println(err)
		return "", err
	}
	return string(body), err
}
func httpRegex(str string) (results []string, err error) {
	//正则表达式，有点菜，只会(.*?)
	regex := "<td align=\"left\">ssr:(.*?)</td>"

	reg := regexp.MustCompile(regex)

	dataS := reg.FindAllSubmatch([]byte(str), -1)

	results = make([]string, 0)

	if dataS != nil {
		for _, v := range dataS {
			results = append(results, string("ssr:"+string(v[1])))
		}
	}
	return results, err
}
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
func RemoveDuplicatesAndEmpty(a []string) (ret []string){
	a_len := len(a)
	for i:=0; i < a_len; i++{
		if (i > 0 && a[i-1] == a[i]) || len(a[i])==0{
			continue;
		}
		ret = append(ret, a[i])
	}
	return
}

