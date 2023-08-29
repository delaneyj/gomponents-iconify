package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 12h4M9 4a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3m6-16a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3"/>`),
		g.Group(children),
	)
}