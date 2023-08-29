package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laundry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8 1L6 3H3S2 3 2 4v9c0 1 1 1 1 1h9s1 0 1-1V2c0-1-1-1-1-1H8zm.5 1h2a.499.499 0 1 1 0 1h-2a.499.499 0 1 1 0-1zm-1 4a3 3 0 1 1 0 6a3 3 0 0 1 0-6z"/>`),
		g.Group(children),
	)
}