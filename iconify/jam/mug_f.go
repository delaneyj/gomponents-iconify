package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MugF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.874 3H16a4 4 0 0 1 4 4v3a4 4 0 0 1-4 4h-1v2a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h7a4.002 4.002 0 0 1 3.874 3zM15 12h1a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-1v7z"/>`),
		g.Group(children),
	)
}