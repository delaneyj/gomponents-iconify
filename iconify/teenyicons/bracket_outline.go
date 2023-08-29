package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BracketOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 7.5h2m13 0h-2m-8 7H3.5a2 2 0 0 1-2-2v-10a2 2 0 0 1 2-2H5m5 14h1.5a2 2 0 0 0 2-2v-10a2 2 0 0 0-2-2H10"/>`),
		g.Group(children),
	)
}