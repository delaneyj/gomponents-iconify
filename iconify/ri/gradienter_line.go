package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GradienterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.049 13H4.06a8.001 8.001 0 0 0 15.877 0h2.012c-.502 5.053-4.765 9-9.95 9c-5.186 0-9.45-3.947-9.951-9Zm0-2c.502-5.053 4.765-9 9.95-9c5.186 0 9.45 3.947 9.951 9h-2.012A8.001 8.001 0 0 0 4.06 11H2.05Zm9.95 3a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}