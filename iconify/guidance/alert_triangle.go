package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 7v8.5m0 1.5v2m11 1.5L22 22s-5.7-.25-10-.25S2 22 2 22l-1-1.5C6 13 11 2 11 2h2s5 11 10 18.5Z"/>`),
		g.Group(children),
	)
}