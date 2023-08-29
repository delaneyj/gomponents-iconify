package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 12a8 8 0 1 1 16 0a8 8 0 0 1-16 0zm8-10C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm0 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm3 2a2 2 0 1 1 4 0a2 2 0 0 1-4 0zm-8-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4z"/></g>`),
		g.Group(children),
	)
}