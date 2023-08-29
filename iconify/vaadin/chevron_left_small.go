package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeftSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m4 8l6.32-6.32L12 3.35L7.35 8L12 12.65l-1.68 1.67L4 8z"/>`),
		g.Group(children),
	)
}