package modules

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v3"
"regexp"
)

var HyperLink = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

func PARSET(c tb.Context) error {

	return c.Reply(ParseMD(c))

}

func ParseMD(c tb.Context) string {
	text := c.Message().ReplyTo.Text
	cor := 0
	for _, x := range c.Message().ReplyTo.Entities {
		offset, length := x.Offset, x.Length
		if x.Type == tb.EntityBold {
			text = string(text[:offset+cor]) + "<b>" + string(text[offset+cor:offset+cor+length]) + "</b>" + string(text[offset+cor+length:])
			cor += 7
		} else if x.Type == tb.EntityCode {
			text = string(text[:offset+cor]) + "<code>" + string(text[offset+cor:offset+cor+length]) + "</code>" + string(text[offset+cor+length:])
			cor += 13
		} else if x.Type == tb.EntityUnderline {
text = string(text[:offset+cor]) + "<code>" + string(text[offset+cor:offset+cor+length]) + "</code>" + string(text[offset+cor+length:])
			cor += 7
} else if x.Type == tb.EntityItalic {
text = string(text[:offset+cor]) + "<code>" + string(text[offset+cor:offset+cor+length]) + "</code>" + string(text[offset+cor+length:])
			cor += 7
}
	}
        Links := HyperLink.FindAllStringSubmatch(text, -1)
	if Links != nil {
		for _, x := range f {
                        if strings.Contains(x[2], "buttonurl") {
continue
}
			text = strings.Replace(text, x[0], fmt.Sprintf("<a href='%s'>%s</a>", x[1], x[2]), 1)
			fmt.Println(text)
		}
	}
	fmt.Println(text)
	return text

}
