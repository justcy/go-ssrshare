package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main(){

	http.HandleFunc("/", ssr)
	http.HandleFunc("/dxxzst", dxxzst)

	http.ListenAndServe(":8922", nil)


	//str := httpGet("https://github.com/dxxzst/Free-SS-SSR")
	//reg := httpRegex(str)
	//result := strings.Join(reg, "\n")
	//strbytes := []byte(result)
	//fmt.Println(base64.StdEncoding.EncodeToString(strbytes))
}

func ssr(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "欢迎使用免费ssr订阅服务")
}

func dxxzst(w http.ResponseWriter, r *http.Request){
	str,err := httpGet("https://github.com/dxxzst/Free-SS-SSR")
	if err != nil{
		fmt.Sprint(w,err)
	}
	reg,err:= httpRegex(str)
	result := strings.Join(reg, "\n")
	strbytes := []byte(result)
	fmt.Fprintf(w, base64.StdEncoding.EncodeToString(strbytes))
}

func httpGet(url string) (str string ,err error){
	resp,err := http.Get(url)
	if err != nil {
		return "",err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		fmt.Println(err)
		return "",err
	}
	return string(body),err
}
func httpRegex(str string) (results []string ,err error) {
	//正则表达式，有点菜，只会(.*?)
	regex := "<td align=\"left\">ssr:(.*?)</td>"

	reg := regexp.MustCompile(regex)

	dataS := reg.FindAllSubmatch([]byte(str), -1)

	results = make([]string,0)

	if dataS != nil{
		for _,v := range dataS {
			results = append(results,string("ssr:"+string(v[1])))
		}
	}
	return results,err
}
