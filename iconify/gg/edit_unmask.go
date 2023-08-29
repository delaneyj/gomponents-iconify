package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditUnmask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-1 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-2 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm10-10a3 3 0 1 1-6 0a3 3 0 0 1 6 0ZM5 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}