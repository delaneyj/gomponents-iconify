package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm-4 2a4 4 0 1 1 7.446 2.032l1.761 1.76a1 1 0 0 1-1.414 1.415l-1.761-1.761A4 4 0 0 1 7 11z"/><path d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16zM2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12z"/></g>`),
		g.Group(children),
	)
}