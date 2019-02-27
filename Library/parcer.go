package Library

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	D "github.com/AntonVTR/AddLableOnPic/Draw"
	"golang.org/x/net/html"
)

type member struct {
	position, name, team, time, group string
}

var memberList []member

func _check(err error) {
	if err != nil {
		panic(err)
	}
}

// основная функция обработки
func parseURL(url string) {
	fmt.Println("request: " + url)

	doc, err := http.Get(url)
	_check(err)
	b := doc.Body
	z := html.NewTokenizer(b)

	group := false
	member := false
	groupsNum := 0
	memberNum := 0
	groupN := ""
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an tag
			if t.Data == "h2" {
				group = true
				continue
			}

		case tt == html.TextToken:
			t := z.Token()
			if group {
				groupN = t.Data
				groupsNum++
				group = false

			} else if member {
				addMembers(t.Data, groupN)
				memberNum++
			}

		case tt == html.EndTagToken:
			t := z.Token()
			if t.Data == "u" {
				member = true
			} else if t.Data == "pre" {
				member = false
			}
		}
	}
}

func addMembers(m string, g string) {

	lines := strings.Split(m, "\n")
	for _, l := range lines {
		m := strings.Split(l, "  ")
		if len(m) > 1 {

			addMember(m, g)
		}

	}
}
func addMember(m []string, g string) {
	i := 0
	mem := new(member)
	mem.group = g
	for _, v := range m {
		if v != "" {
			switch i {
			case 0:
				mem.position = strings.TrimFunc(v, func(r rune) bool { return !unicode.IsNumber(r) })
				mem.name = strings.TrimFunc(v, func(r rune) bool { return !unicode.IsLetter(r) })
				i++
			case 1:
				mem.team = strings.TrimFunc(v, func(r rune) bool { return !unicode.IsLetter(r) })
				if mem.team == "" {
					re := regexp.MustCompile(`\d\d:\d\d:\d\d,\d`)
					mem.time = string(re.Find([]byte(v)))

				}
				i++
			case 2:
				re := regexp.MustCompile(`\d\d:\d\d:\d\d,\d`)
				mem.time = string(re.Find([]byte(v)))
				i = 33
			}
		}

	}
	//fmt.Println(positionM, "name", nameM, "team", teamM, "time", timeM, "group", g)
	//mem := member{positionM, nameM, teamM, timeM, g}
	if mem.time != "" {
		memberList = append(memberList, *mem)
	}
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

//ckmment dfedf
func ParseU(url string) {

	fmt.Println("Start")
	parseURL(url)
	for _, mem := range memberList {
		fmt.Println(mem)
		D.AddText(mem.name, mem.group, mem.time, mem.position)
	}

}
