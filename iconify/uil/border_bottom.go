package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13.5a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0-8a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-4-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0 8a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm12-8a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-4 8a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm-4-8a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm4 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm4 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 8H4a1 1 0 0 0 0 2h16a1 1 0 0 0 0-2Zm0-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-16 6a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0 8a1 1 0 1 0-1-1a1 1 0 0 0 1 1Zm0-12a1 1 0 1 0-1-1a1 1 0 0 0 1 1Z"/>`),
		g.Group(children),
	)
}