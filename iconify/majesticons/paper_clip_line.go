package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperClipLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M18 6a1 1 0 0 1 1 1v8a7 7 0 1 1-14 0V7a5 5 0 0 1 10 0v8a3 3 0 1 1-6 0V7a1 1 0 1 1 2 0v8a1 1 0 1 0 2 0V7a3 3 0 1 0-6 0v8a5 5 0 0 0 10 0V7a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}