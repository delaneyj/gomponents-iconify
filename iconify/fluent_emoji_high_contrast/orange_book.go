package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.5 3.5a2 2 0 0 1 2-2H25A2.5 2.5 0 0 1 27.5 4v23.5H7a.5.5 0 0 0 0 1h20.622l-.236.667A2 2 0 0 1 25.5 30.5h-19a2 2 0 0 1-2-2v-25Zm20.645 22l1.355-1.355v-3.586L21.559 25.5h3.586ZM8.5 3.559v3.586L13.145 2.5H9.559L8.5 3.559ZM14.559 2.5L8.5 8.559v3.586L18.145 2.5h-3.586Zm5 0L8.5 13.559v3.586L23.145 2.5h-3.586Zm5 0L8.5 18.559v3.586l18-18V4A1.5 1.5 0 0 0 25 2.5h-.441ZM26.5 5.559l-18 18V25.5h1.645L26.5 9.145V5.559Zm0 5L11.559 25.5h3.586L26.5 14.145v-3.586Zm0 5L16.559 25.5h3.586l6.355-6.355v-3.586Z"/>`),
		g.Group(children),
	)
}