package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZAxis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 22v3.586l-9-9V5.828l2.586 2.586L21 7l-5-5l-5 5l1.414 1.414L15 5.828v10.758l-9 9V22H4v7h7v-2H7.414L16 18.414L24.586 27H21v2h7v-7h-2z"/>`),
		g.Group(children),
	)
}