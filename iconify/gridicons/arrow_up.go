package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20V7.83l5.59 5.59L20 12l-8-8l-8 8l1.41 1.41L11 7.83V20h2z"/>`),
		g.Group(children),
	)
}