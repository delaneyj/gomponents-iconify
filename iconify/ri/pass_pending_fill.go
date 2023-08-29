package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PassPendingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h20a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H2Zm7 6a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm2 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-5.473 7.025l-1.414-1.414A5.486 5.486 0 0 1 8.003 14c1.518 0 2.894.617 3.888 1.61l-1.414 1.415A3.486 3.486 0 0 0 8.002 16c-.967 0-1.84.39-2.475 1.025ZM13 15V9h2v6h-2Zm4 0V9h2v6h-2Z"/>`),
		g.Group(children),
	)
}