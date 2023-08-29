package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionRotaryStraight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 12.09V5.828l4.586 4.586L23 9l-7-7l-7 7l1.414 1.414L15 5.828v6.262a5.992 5.992 0 0 0 0 11.82V28h2v-4.09a5.992 5.992 0 0 0 0-11.82ZM16 22a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}