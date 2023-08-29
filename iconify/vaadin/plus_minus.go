package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 7h6v2h-6V7zM4 5H2v2H0v2h2v2h2V9h2V7H4zm2-3l3 12h1L7 2z"/>`),
		g.Group(children),
	)
}