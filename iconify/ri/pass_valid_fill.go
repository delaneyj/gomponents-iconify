package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassValidFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h20a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H2Zm7 6a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm2 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-5.473 7.025l-1.414-1.414A5.486 5.486 0 0 1 8.003 14c1.518 0 2.894.617 3.888 1.61l-1.414 1.415A3.486 3.486 0 0 0 8.002 16c-.967 0-1.84.39-2.475 1.025Zm14.68-6.318l-4 4l-.707.707l-.707-.707l-2.5-2.5l1.414-1.414l1.793 1.793l3.293-3.293l1.414 1.414Z"/>`),
		g.Group(children),
	)
}