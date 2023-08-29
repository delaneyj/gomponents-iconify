package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubleSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M11 6v10H9v2h2v2H9v2h2v4h2v-4h5v-2h-5v-2h6c3.3 0 6-2.7 6-6s-2.7-6-6-6zm2 2h6c2.219 0 4 1.781 4 4c0 2.219-1.781 4-4 4h-6z"/>`),
		g.Group(children),
	)
}