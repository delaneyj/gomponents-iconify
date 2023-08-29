package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 18.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm9-15a1 1 0 0 0-1-1h-16a1 1 0 0 0-1 1v16a1 1 0 0 0 2 0v-15h15a1 1 0 0 0 1-1Zm-5 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm12-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}