package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 23.3l-.6-.5c-2-1.9-3.4-3.1-3.4-4.6c0-1.2 1-2.2 2.2-2.2c.7 0 1.4.3 1.8.8c.4-.5 1.1-.8 1.8-.8c1.2 0 2.2.9 2.2 2.2c0 1.5-1.4 2.7-3.4 4.6l-.6.5M6 22a2 2 0 0 1-2-2V4c0-1.11.89-2 2-2h1v7l2.5-1.5L12 9V2h6a2 2 0 0 1 2 2v9.08L19 13a6.005 6.005 0 0 0-5.2 9H6Z"/>`),
		g.Group(children),
	)
}