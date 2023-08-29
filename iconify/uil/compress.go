package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 9h5a1 1 0 0 0 0-2h-4V3a1 1 0 0 0-2 0v5a1 1 0 0 0 1 1Zm-8 6H3a1 1 0 0 0 0 2h4v4a1 1 0 0 0 2 0v-5a1 1 0 0 0-1-1ZM8 2a1 1 0 0 0-1 1v4H3a1 1 0 0 0 0 2h5a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Zm13 13h-5a1 1 0 0 0-1 1v5a1 1 0 0 0 2 0v-4h4a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}