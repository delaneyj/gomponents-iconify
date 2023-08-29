package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Album(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="11.99" cy="11.99" r="2.01" fill="currentColor"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8z"/><path fill="currentColor" d="M12 6a6 6 0 0 0-6 6h2a4 4 0 0 1 4-4z"/>`),
		g.Group(children),
	)
}