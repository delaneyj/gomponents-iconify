package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Man(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M9 1.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 9 1.5zM9 4H6a1 1 0 0 0-1 1v5h1v6h1.25v-6h.5v6H9v-6h1V5a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}