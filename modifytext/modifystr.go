package modifytext

import (
	"fmt"
	"regexp"
	"strings"
)

/*
represents an issue.
reg is a regex for searching the issue,
replaceWith is the function that returns the string, which the matched string will be replaced with
*/
type issue struct {
	reg         *regexp.Regexp
	replaceWith func([]string) (string, error)
}

// map for processed issues
var issues = map[string]*issue{
	"hex": {
		reg:         regexp.MustCompile(`(\b\w+\b) \(hex\)`),
		replaceWith: func(s []string) (string, error) { return convertHexToDec(s[1]) },
	},
	"bin": {
		reg:         regexp.MustCompile(`(\b\w+\b) \(bin\)`),
		replaceWith: func(s []string) (string, error) { return convertBinToDec(s[1]) },
	},
	"up": {
		reg:         regexp.MustCompile(`(\b\w+\b) \(up\)`),
		replaceWith: func(s []string) (string, error) { return strings.ToUpper(s[1]), nil },
	},
	"low": {
		reg:         regexp.MustCompile(`(\b\w+\b) \(low\)`),
		replaceWith: func(s []string) (string, error) { return strings.ToLower(s[1]), nil },
	},
	"cap": {
		reg:         regexp.MustCompile(`(\b\w+\b) \(cap\)`),
		replaceWith: func(s []string) (string, error) { return Capitalize(s[1]), nil },
	},
	"punctuation": {
		reg: regexp.MustCompile(`([\w\)'])(\s)+(\.\.\.|!\?|\.|\,|!|\?|;|:)(\s*)([\(\w'])?`),
		replaceWith: func(s []string) (string, error) {
			res := s[1] + s[3]
			if s[4] == "" && s[5] != "" {
				res += " " + s[5]
			} else { // keep all spaces
				res += s[4] + s[5]
			}
			return res, nil
		},
	},
	"'": {
		reg:         regexp.MustCompile(`'(\s*)([^']*?)(\s*)'`),
		replaceWith: func(s []string) (string, error) { return "'" + s[2] + "'", nil },
	},
	"an": {
		reg:         regexp.MustCompile(`([aA])(\W+[AaEeUuIiOoHh])`),
		replaceWith: func(s []string) (string, error) { return s[1] + "n" + s[2], nil },
	},
}

// order for fixing issues
var order = []string{"hex", "bin", "up", "low", "cap", "punctuation", "'", "an"}

/*
modifies the given string according to list in README.md
*/
func ModifyString(str string) (string, error) {
	var err, er error

	// fix all issue except for the ones with numbers  (low, <number>, etc.)
	for _, name := range order {
		str, er = issues[name].solveIssue(str)
		if err != nil {
			err = fmt.Errorf("the string consists something wrong. The last error occurs when solved the issue %s. %s", name, er)
		}
	}

	// look for (low), (up), (cap) with a number after it, like so: (low, <number>)
	matches := regexp.MustCompile(`\((low|up|cap), (\d+)\)`).FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		// create new issue for the concrete number after low|up|cap and resolve it.
		is := issue{
			reg:         regexp.MustCompile(`((?:\b\w+\b\W*){` + string(match[2]) + `}) \(` + match[1] + `, ` + match[2] + `\)`),
			replaceWith: issues[match[1]].replaceWith,
		}
		str, er = is.solveIssue(str)
		if err != nil {
			err = fmt.Errorf("the string consists something wrong. The last error occurs when solved the issue %s. %s", fmt.Sprintf("(%s, %s)", match[1], match[2]), er)
		}
	}
	return str, err
}

/*
looks for the issue in the given string and returns a fixed string
*/
func (is *issue) solveIssue(str string) (string, error) {
	var err error
	matches := is.reg.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		// new is a string which will have used for the strings.Replace function
		new, er := is.replaceWith(match)
		if er != nil {
			err = fmt.Errorf("there were at least 1 error. The last occures with the match %s: %s", match[0], er)
		}
		str = strings.Replace(str, match[0], new, 1)
	}
	return str, err
}
