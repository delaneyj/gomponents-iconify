package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boardgamegeek(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.7 4.44l-2.38.64L19.65 0L4.53 5.56l.83 6.67l-1.4 1.34L8.12 24l8.85-3.26l3.07-7.22l-1.32-1.27l.98-7.81Z"/>`),
		g.Group(children),
	)
}