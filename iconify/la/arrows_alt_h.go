package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsAltH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m9.281 6.781l-8.5 8.5L.094 16l.687.719l8.5 8.5l1.438-1.438L3.938 17h24.125l-6.782 6.781l1.438 1.438l8.5-8.5l.687-.719l-.687-.719l-8.5-8.5L21.28 8.22L28.063 15H3.938l6.78-6.781z"/>`),
		g.Group(children),
	)
}