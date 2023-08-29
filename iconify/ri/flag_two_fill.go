package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h19.138a.5.5 0 0 1 .434.748L18 10l3.573 6.252a.5.5 0 0 1-.435.748H4v5H2V3Z"/>`),
		g.Group(children),
	)
}