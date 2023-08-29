package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12zm5.5 1.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zm3-1.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0zm4.5 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}