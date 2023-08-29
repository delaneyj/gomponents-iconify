package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompressLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.979 9.457l-3.57-.003l-.002 2l7 .006l.006-7l-2-.002L9.41 8.06L3.096 1.77L1.685 3.185l6.294 6.271Zm11.582 5.095l-.008-2l-7 .028l.028 7l2-.008l-.014-3.601l6.343 6.26l1.405-1.424l-6.324-6.24l3.57-.015Z"/>`),
		g.Group(children),
	)
}