package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M374.2 189.2H157.5L275.7 71.1H157.5L0 228.6l157.5 157.5h118.2L157.5 268h216.6c32.6 0 59.1 26.4 59.1 59.1c0 32.6-26.4 59.1-59.1 59.1V465c76.1 0 137.8-61.7 137.8-137.8c.1-76.3-61.6-138-137.7-138z"/>`),
		g.Group(children),
	)
}