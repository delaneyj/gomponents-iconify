package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDetailsLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 6h9v2H8zm0-3h11v2H8zM1 3h6v6H1zm7 11h9v2H8zm0-3h11v2H8zm-7 0h6v6H1z"/>`),
		g.Group(children),
	)
}