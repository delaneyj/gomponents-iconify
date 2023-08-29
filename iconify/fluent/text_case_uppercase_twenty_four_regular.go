package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextCaseUppercaseTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.707 3.5a.75.75 0 0 0-1.406-.02l-6 15.5a.75.75 0 1 0 1.399.54l1.556-4.02h7.118l1.42 4a.75.75 0 1 0 1.413-.5l-5.5-15.5ZM3.837 14L6.97 5.907L9.84 14H3.837ZM14.75 3.25A.75.75 0 0 0 14 4v15.25c0 .414.336.75.75.75h4.125a4.625 4.625 0 0 0 2.006-8.793A4.5 4.5 0 0 0 18 3.25h-3.25ZM21 7.75a3 3 0 0 1-3 3h-2.5v-6H18a3 3 0 0 1 3 3ZM18.875 18.5H15.5v-6.25h3.375a3.125 3.125 0 1 1 0 6.25Z"/>`),
		g.Group(children),
	)
}