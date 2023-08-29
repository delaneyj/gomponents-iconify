package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuHotdog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6a3 3 0 0 0-3 3h16a3 3 0 0 0-3-3H7Zm0 12a3 3 0 0 1-3-3h16a3 3 0 0 1-3 3H7Zm-4-7a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H3Z"/>`),
		g.Group(children),
	)
}