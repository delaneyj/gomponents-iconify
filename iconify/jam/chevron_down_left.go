package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 8h7a1 1 0 1 1 0 2H1a1 1 0 0 1-1-1V1a1 1 0 1 1 2 0v7z"/>`),
		g.Group(children),
	)
}