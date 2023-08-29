package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuOreos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3a3 3 0 0 0-3 3h16a3 3 0 0 0-3-3H7Zm0 8a3 3 0 0 1-3-3h16a3 3 0 0 1-3 3H7Zm0 2a3 3 0 0 0-3 3h16a3 3 0 0 0-3-3H7Zm0 8a3 3 0 0 1-3-3h16a3 3 0 0 1-3 3H7Z"/>`),
		g.Group(children),
	)
}