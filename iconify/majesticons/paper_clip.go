package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperClip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M6 2a4 4 0 0 0-4 4v12a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4H6zm13 7a1 1 0 1 0-2 0v4a5 5 0 0 1-10 0V9a3 3 0 0 1 6 0v4a1 1 0 1 1-2 0V9a1 1 0 1 0-2 0v4a3 3 0 1 0 6 0V9A5 5 0 0 0 5 9v4a7 7 0 1 0 14 0V9z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}