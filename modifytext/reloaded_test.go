package modifytext

import (
	"regexp"
	"testing"
)

// TestHellostr calls greetings.Hello with a str, checking
// for a valid return value.

func TestModifyString(t *testing.T) {
	str := []string{
		"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
		"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
		"Don not be sad ,because sad backwards is das . And das not good",
		"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
		"Are you ready!? I was thinking ... You were right ,werent you !?",
		"If a I make ...you BREAKFAST , IN BED (low, 3) just say thank you .() instead of (cap) : how (cap) did you get in my house (up, 2) ?",
		"A elephant is not (up) a ' horse' !",
	}
	want := []string{
		"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		"I have to pack 5 outfits. Packed 26 just to be sure",
		"Don not be sad, because sad backwards is das. And das not good",
		"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		"Are you ready!? I was thinking... You were right, werent you!?",
		"If an I make... you breakfast, in bed just say thank you. () instead Of: How did you get in MY HOUSE?",
		"An elephant is NOT an 'horse'!",
	}
	for i := 0; i < len(str); i++ {
		res, err := ModifyString(str[i])
		if res != want[i] || err != nil {
			t.Fatalf("\nModifyString \n(%s) \n=       %s, %v, \nwant to %s, <nil>\n", str[i], res, err, want[i])
		}
	}
}

func TestModifyString1(t *testing.T) {
	str := []string{
		"If I make you BREAKFAST IN BED (low) just say thank you instead of: how did you get in my house (up) ?",
	}
	want := []string{
		"If I make you BREAKFAST IN bed just say thank you instead of: how did you get in my HOUSE?",
	}
	for i := 0; i < len(str); i++ {
		res, err := ModifyString(str[i])
		if res != want[i] || err != nil {
			t.Fatalf("\nModifyString \n(%s) \n=       %s, %v, \nwant to %s, nil\n", str[i], res, err, want[i])
		}
	}
}

func BenchmarkRegex(t *testing.B) {
	str := []string{
		" fgfgh ?",
	}
	res := regexp.MustCompile(`(\w)?(\s*)(\.\.\.|!\?|\.|\,|!|\?|;|:)(\s)*(\w)?`).FindAllStringSubmatch(str[0], -1)

	for m, match := range res {
		t.Logf("match %d: [%s]\n", m, match[0])
		for s, submatch := range match[1:] {
			t.Logf("submatch %d: [%s] ", s, submatch)
		}
	}
}
