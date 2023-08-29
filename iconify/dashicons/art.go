package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Art(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.55 3.06c1.01.34-1.95 2.01-.1 3.13c1.04.63 3.31-2.22 4.45-2.86c.97-.54 2.67-.65 3.53 1.23c1.09 2.38.14 8.57-3.79 11.06c-3.97 2.5-8.97 1.23-10.7-2.66c-2.01-4.53 3.12-11.09 6.61-9.9zm1.21 6.45c.73 1.64 4.7-.5 3.79-2.8c-.59-1.49-4.48 1.25-3.79 2.8z"/>`),
		g.Group(children),
	)
}